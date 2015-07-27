var pezPortal = angular.module('pezPortal', [], function($interpolateProvider) {
      $interpolateProvider.startSymbol('{*{');
      $interpolateProvider.endSymbol('}*}');
  })
  .controller('PezPortalController', function($scope, $http, $timeout, $window) {
    $scope.hideCLIExample = true;
    var myData = {};
    var pauth = this;
    var restUriBase = "/v1/auth/api-key";
    var restOrgUriBase = "/v1/org/user";
    var meUri = "/me";
    var messaging = {
      "hasOrgBtn": "View Org Now",
      "createOrgBtn": "Create Your Org Now",
      "noApiKey": "You don't have a key yet",
      "loading": "Loading... Please Wait",
      "oktaSetup": "Get Okta Tile for HeritageCF",
      "invalidUser": "query failed. unable to find matching user guid."
    };

    var urls = {
      "okta": "http://login.run.pez.pivotal.io/saml/login/alias/login.run.pez.pivotal.io?disco=true",
      "oktaHome": "https://pivotal.okta.com/app/UserHome"
    };

    $timeout(function () {
      callMeUsingVerb($http.get, meUri);
    }, 1);

    pauth.getRestUri = function() {
      return [restUriBase, $scope.myEmail].join("/");
    }

    pauth.getOrgRestUri = function() {
      return [restOrgUriBase, $scope.myEmail].join("/");
    }

    pauth.getorg = function() {
      console.log(pauth.getOrgRestUri());
      getOrgStatus(pauth.getOrgRestUri());
    };

    pauth.createorg = function() {

      if ($scope.orgButtonText === messaging.createOrgBtn) {
        createOrg(pauth.getOrgRestUri());

      } else if ($scope.orgButtonText === messaging.hasOrgBtn) {
        $window.location.href = urls.okta;

      }  else if ($scope.orgButtonText === messaging.oktaSetup) {
        $window.location.href = urls.oktaHome;
      }
      $scope.orgButtonText = messaging.loading;
    };

    pauth.create = function() {
      callAPIUsingVerb($http.put, pauth.getRestUri());
    };

    pauth.remove = function() {
      callAPIUsingVerb($http.delete, pauth.getRestUri());
    };

    function callMeUsingVerb(verbCaller, uri) {
      var responsePromise = verbCaller(uri);
      responsePromise.success(function(data, status, headers, config) {
          $scope.myName = data.Payload.displayName;
          $scope.myEmail = data.Payload.emails[0].value;
          $scope.displayName = $scope.myName ? $scope.myName : $scope.myEmail
          callAPIUsingVerb($http.get, pauth.getRestUri());
          pauth.getorg();
      });
    }

    function createOrg(uri) {
      var responsePromise = $http.put(uri);
      responsePromise.success(function(data, status, headers, config) {
        console.log(data);
        $scope.orgButtonText = messaging.hasOrgBtn;
        $scope.hideCLIExample = false;
      });

      responsePromise.error(function(data, status, headers, config) {
          var forwardToOkta = false;

          if(status === 403) {
            console.log(data.ErrorMsg);

            if (messaging.invalidUser == data.ErrorMsg) {
              forwardToOkta = confirm("You have not set up your account in Okta yet. Please head over to Okta and click on the `PEZ HeritageCF` tile.");
            }

            if ( forwardToOkta === true) {
              $window.location.href = urls.oktaHome;

            } else {
              $scope.orgButtonText = messaging.oktaSetup;
            }
          }
      });
    }

    function getOrgStatus(uri) {
      var responsePromise = $http.get(uri);
      responsePromise.success(function(data, status, headers, config) {
        console.log(data);
        $scope.orgButtonText = messaging.hasOrgBtn;
        $scope.hideCLIExample = false;
      });

      responsePromise.error(function(data, status, headers, config) {

          if(status === 403) {
            $scope.orgButtonText = messaging.createOrgBtn;
            console.log(data.ErrorMsg);
          }
      });
    }

    function callAPIUsingVerb(verbCaller, uri) {
      var responsePromise = verbCaller(uri);

      responsePromise.success(function(data, status, headers, config) {
          $scope.myData = data;
          $scope.myApiKey = data.APIKey;
      });

      responsePromise.error(function(data, status, headers, config) {
        $scope.myApiKey = messaging.noApiKey;
        pauth.create();
      });
    }
  });
