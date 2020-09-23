package michttp

import "sync"

type HttpResponseMap map[string]interface{}
type HttpRequestMap map[string]MicHttpRequest
type Response struct {
	KeyName string
	Err error
	Info string
}

func RequesAll(requestMap HttpRequestMap) HttpResponseMap  {
	req :=HttpResponseMap{}
	var wg sync.WaitGroup
	for key,mic :=range requestMap{
		wg.Add(1)
		go Http(key,mic,req,&wg)
	}
	wg.Wait()
	return req
}

func Http(keyName string,request MicHttpRequest,returnMap HttpResponseMap,wg *sync.WaitGroup  )  {
	defer wg.Done()
	info,err:=request.GetRequesString()
	returnMap[keyName]=Response{
		Err:err,
		Info: info,
		KeyName: keyName,
	}
}








