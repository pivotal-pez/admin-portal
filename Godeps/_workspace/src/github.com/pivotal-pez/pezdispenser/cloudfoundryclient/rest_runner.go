package cloudfoundryclient

import "net/http"

func (s *RestRunner) Run() {
	var (
		req *http.Request
		res *http.Response
		err error
	)
	s.Logger.Println("making rest call to: url:", s.URL, "- with verb:", s.Verb, "- using data:", s.Data, "- at path:", s.Path)

	if _, err = s.RequestDecorator.Login(); err != nil {
		s.OnFailure(res, err)
	}

	if req, err = s.RequestDecorator.CreateAuthRequest(s.Verb, s.URL, s.Path, s.Data); err == nil {
		s.Logger.Println("we created the decorated request")

		if s.Verb == "GET" && s.Data != nil {
			switch v := s.Data.(type) {
			case string:
				req.URL.RawQuery = v
			}
		}

		if res, err = s.RequestDecorator.HttpClient().Do(req); res.StatusCode == s.SuccessStatusCode && err == nil {
			defer res.Body.Close()
			s.Logger.Println("we are now going to execute the success callback")
			s.OnSuccess(res)

		} else {
			s.Logger.Println("we are now going to execute the failure callback")
			s.OnFailure(res, getErrorCode(err))
		}
	}
}

func getErrorCode(err error) error {
	if err == nil {
		err = ErrFailedStatusCode
	}
	return err
}
