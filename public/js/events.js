angular.module('EventsApp', [])
  .controller('EventsController', function($scope, $http, $interval, $timeout, $window) {
    var events = this;
    var refreshInterval = 5000;
    var eventsInfoURI = "/v1/info/events";
    var messaging = {
      unavailable: "** Service Unavailable **",
      empty: "-",
      zero: 0
    };

    $interval( function(){ events.getEventsInfo(); }, refreshInterval, 0, true);

    $timeout(function () {
      events.getEventsInfo();
    }, 1);

    events.getEventsInfo = function() {
      var responsePromise = $http.get(eventsInfoURI);
      responsePromise.success(function(data, status, headers, config) {
        $scope.events = data;
        $scope.userInfoStatus = "";
      });

      responsePromise.error(function(data, status, headers, config) {
        $scope.events = {};
        $scope.userInfoStatus = messaging.unavailable;
        console.log(status, data);
      });
    };
  });
