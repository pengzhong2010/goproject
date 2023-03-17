package es7

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

type EsClient struct {
	urls []string
	Conn *elastic.Client
}

func NewEs7(urls []string, username, password string) (db *EsClient, cf func(), err error) {
	db = &EsClient{
		urls: urls,
	}
	db.Conn, err = elastic.NewClient(
		elastic.SetURL(urls...),
		elastic.SetBasicAuth(username, password),
		elastic.SetSniff(false),
	)
	if err != nil {
		return
	}
	cf = func() {}
	return
}

func (t *EsClient) Ping(c context.Context) (out string, err error) {
	info, code, err := t.Conn.Ping(t.urls[0]).Do(c)
	if err != nil {
		return
	}
	out = fmt.Sprintf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	return
}
func (t *EsClient) ExistIndex(c context.Context, indexName string) (exists bool, err error) {
	exists, err = t.Conn.IndexExists(indexName).Do(c)
	return
}

func (t *EsClient) CreateIndex(c context.Context, indexName string, mapping string) (err error) {
	var exists bool
	exists, err = t.ExistIndex(c, indexName)
	if err != nil {
		err = errors.Wrap(err, "IndexExists err")
		return
	}
	if exists {
		err = errors.Wrap(err, "IndexExists exist")
		return
	}
	createIndex, errE1 := t.Conn.CreateIndex(indexName).Body(mapping).Do(c)
	if errE1 != nil {
		err = errors.Wrap(errE1, "CreateIndex err")
		return
	}
	if !createIndex.Acknowledged {
		err = errors.Wrap(nil, "Acknowledged err")
		return
	}
	return
}

func (t *EsClient) DeleteIndex(c context.Context, indexName string) (err error) {
	exists, err := t.ExistIndex(c, indexName)
	if err != nil {
		err = errors.Wrap(err, "IndexExists err")
		return
	}
	if !exists {
		err = errors.Wrap(err, "IndexExists not exist")
		return
	}
	deleteIndex, err := t.Conn.DeleteIndex(indexName).Do(c)
	if err != nil {
		err = errors.Wrap(err, "DeleteIndex err")
		return
	}
	if !deleteIndex.Acknowledged {
		err = errors.Wrap(err, "Acknowledged err")
		return
	}
	return
}

func (t *EsClient) AddData(c context.Context, indexName string, id int64, data interface{}) (err error) {
	_, err = t.Conn.Index().Index(indexName).Id(strconv.FormatInt(id, 10)).BodyJson(data).Do(c)
	if err != nil {
		err = errors.Wrap(err, "AddData err")
		return
	}
	return
}

func (t *EsClient) AddDataBulk(c context.Context, indexName string, data map[int64]interface{}) (err error) {
	bulkReq := t.Conn.Bulk()
	for id, i := range data {
		tmp := i
		//req := elastic.NewBulkIndexRequest().Index(indexName).Doc(tmp)
		req := elastic.NewBulkCreateRequest().Index(indexName).Id(strconv.FormatInt(id, 10)).Doc(tmp)
		bulkReq = bulkReq.Add(req)
	}
	_, err = bulkReq.Do(c)
	if err != nil {
		err = errors.Wrap(err, "AddDataBulk err")
		return
	}
	return
}
func (t *EsClient) UpdateData(c context.Context, indexName string, id int64, data interface{}) (err error) {
	_, err = t.Conn.Update().Index(indexName).Id(strconv.FormatInt(id, 10)).Doc(data).Do(c)
	if err != nil {
		err = errors.Wrap(err, "UpdateData err")
		return
	}
	return
}
func (t *EsClient) UpdateDataBulk(c context.Context, indexName string, data map[int64]interface{}) (err error) {
	bulkReq := t.Conn.Bulk()
	for id, i := range data {
		tmp := i
		req := elastic.NewBulkUpdateRequest().Index(indexName).Id(strconv.FormatInt(id, 10)).Doc(tmp).DocAsUpsert(true)
		bulkReq = bulkReq.Add(req)
	}
	_, err = bulkReq.Do(c)
	if err != nil {
		err = errors.Wrap(err, "UpdateDataBulk err")
		return
	}
	return
}

func (t *EsClient) Refresh(c context.Context, indexName string) (err error) {
	_, err = t.Conn.Refresh().Index(indexName).Do(c)
	if err != nil {
		err = errors.Wrap(err, "Refresh err")
		return
	}
	return
}

func (t *EsClient) AddAliasIndex(c context.Context, indexName, aliasName string) (err error) {
	exists, err := t.ExistIndex(c, indexName)
	if err != nil {
		err = errors.Wrap(err, "IndexExists err")
		return
	}
	if !exists {
		err = errors.Wrap(err, "IndexExists not exist")
		return
	}
	res, errE1 := t.Conn.Alias().Add(indexName, aliasName).Do(c)
	if errE1 != nil {
		err = errors.Wrap(errE1, "Add AliasIndex err")
		return
	}
	if !res.Acknowledged {
		err = errors.Wrap(nil, "Acknowledged err")
		return
	}
	return
}
func (t *EsClient) RemoveAliasIndex(c context.Context, indexName, aliasName string) (err error) {
	exists, err := t.ExistIndex(c, indexName)
	if err != nil {
		err = errors.Wrap(err, "IndexExists err")
		return
	}
	if !exists {
		err = errors.Wrap(err, "IndexExists not exist")
		return
	}
	//log.Info("indexName:%+v", indexName)
	//log.Info("aliasName:%+v", aliasName)
	res, errE1 := t.Conn.Alias().Remove(indexName, aliasName).Do(c)
	if errE1 != nil {
		err = errors.Wrap(errE1, "Remove AliasIndex err")
		return
	}
	if !res.Acknowledged {
		err = errors.Wrap(nil, "Acknowledged err")
		return
	}
	return
}
func (t *EsClient) GetDocById(c context.Context, indexName string, id int64) (res json.RawMessage, err error) {
	// Get tweet with specified ID
	doc, errT := t.Conn.Get().
		Index(indexName).
		Id(strconv.FormatInt(id, 10)).
		Do(c)
	if errT != nil {
		switch {
		case elastic.IsNotFound(errT):
			err = errors.Wrap(errT, "Document not found")
			return
		case elastic.IsTimeout(errT):
			err = errors.Wrap(errT, "Timeout retrieving document")
			return
		case elastic.IsConnErr(errT):
			err = errors.Wrap(errT, "Connection problem")
			return
		default:
			err = errors.Wrap(errT, "Es query err")
			return
		}
	}
	if !doc.Found {
		return
	}
	res = doc.Source
	return
}

// func to print query dsl
func (t *EsClient) PrintQuery(src interface{}) {
	fmt.Println("*****")
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println("*****")
	return
}

//
//// query dsl demo
//func (t *EsClient) printQueryDemo() {
//	query := elastic.NewTermQuery("genres", "动画")
//	src, err := query.Source()
//	if err != nil {
//		panic(err)
//	}
//	Client.PrintQuery(src)
//	return
//}
//func Test() (err error) {
//	fmt.Println("----------------------")
//	indexName := "incident-20220303120939"
//
//	tmp := Client
//	ctx := context.Background()
//	err = tmp.DeleteIndex(ctx, indexName)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	indexMapping := make(map[string]string)
//	indexMapping[indexName] = `
//{
//		"settings":{
//			"number_of_shards": 1,
//			"number_of_replicas": 0,
//			"refresh_interval": -1
//		},
//		"mappings":{
//			"properties":{
//				"user":{
//					"type":"keyword"
//				},
//				"message":{
//					"type":"text",
//					"store": true,
//					"fielddata": true
//				},
//				"image":{
//					"type":"keyword"
//				},
//				"created":{
//					"type":"date"
//				},
//				"tags":{
//					"type":"keyword"
//				},
//				"location":{
//					"type":"geo_point"
//				},
//				"suggest_field":{
//					"type":"completion"
//				}
//			}
//		}
//	}
//`
//	indexMapping[indexName] = `
//{
//		"settings":{
//			"number_of_shards": 1,
//			"number_of_replicas": 0,
//			"refresh_interval": -1
//		},
//		"mappings":{
//			"properties":{
//				"item_id":{
//					"type":"keyword"
//				},
//				"uuid":{
//					"type":"keyword"
//				},
//				"title":{
//					"type":"text"
//				},
//				"incident_type":{
//					"type":"integer"
//				},
//				"incident_status":{
//					"type":"integer"
//				},
//				"creator_uid":{
//					"type":"long"
//				},
//				"creator_uname":{
//					"type":"text"
//				},
//				"ctime":{
//					"type":"long"
//				},
//				"mtime":{
//					"type":"long"
//				},
//				"is_deleted":{
//					"type":"long"
//				},
//				"scene_id":{
//					"type":"long"
//				},
//				"scene_name":{
//					"type":"text"
//				},
//				"start_time_unix":{
//					"type":"long"
//				},
//				"end_time_unix":{
//					"type":"long"
//				},
//				"start_time_plan_unix":{
//					"type":"long"
//				},
//				"during_time_plan_unix":{
//					"type":"long"
//				},
//				"end_time_plan_unix":{
//					"type":"long"
//				},
//				"operator_uids":{
//					"type":"long"
//				},
//				"operator_unames":{
//					"type":"text"
//				},
//				"approver_uids":{
//					"type":"long"
//				},
//				"approver_unames":{
//					"type":"text"
//				},
//				"detail":{
//					"type":"text"
//				},
//				"effect":{
//					"type":"text"
//				},
//				"component_ids":{
//					"type":"long"
//				},
//				"component_names":{
//					"type":"text"
//				},
//				"component_effect_ids":{
//					"type":"long"
//				},
//				"component_effect_names":{
//					"type":"text"
//				},
//				"subscription_uids":{
//					"type":"long"
//				},
//				"subscription_unames":{
//					"type":"text"
//				},
//				"subscription_rids":{
//					"type":"long"
//				},
//				"subscription_rnames":{
//					"type":"text"
//				}
//			}
//		}
//	}
//`
//	err = tmp.CreateIndex(ctx, indexName, indexMapping[indexName])
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("----------------------")
//	return
//}
//
//func Test2() (err error) {
//	fmt.Println("----------------------t2")
//	indexName := "incident-t1"
//	indexNameAlias := "incident-t2"
//	tmp := Client
//	ctx := context.Background()
//
//	err = tmp.AddAliasIndex(ctx, indexName, indexNameAlias)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//err = tmp.RemoveAliasIndex(ctx, indexName, indexNameAlias)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	fmt.Println("----------------------")
//	return
//}
//func Test3() (err error) {
//	b, err := Client.ExistIndex(context.Background(), "incident-20220303120401")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Print(b)
//	return
//}
