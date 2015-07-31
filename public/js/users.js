angular.module('UsersApp', [])
  .controller('UsersController', function($scope, $http, $interval, $timeout, $window) {
    var users = this;
    var refreshInterval = 30000;
    var userInfoURI = "/v1/info/users";
    var messaging = {
      unavailable: "** Service Unavailable **",
      empty: "-"
    };

    $interval( function(){ users.getUsersInfo(); }, refreshInterval, 0, true);

    $timeout(function () {
      users.getUsersInfo();
    }, 1);

    users.getUsersInfo = function() {
      var responsePromise = $http.get(userInfoURI);
      responsePromise.success(function(data, status, headers, config) {
        $scope.totalUsers = data.UserCount;
        $scope.managedUsers = data.ExternalCount;
        $scope.unManagedUsers = data.UAACount;
        $scope.orphanedUsers = data.OrphanedCount;
        $scope.userInfoStatus = "";
      });

      responsePromise.error(function(data, status, headers, config) {
        $scope.totalUsers = messaging.empty;
        $scope.managedUsers = messaging.empty;
        $scope.unManagedUsers = messaging.empty;
        $scope.orphanedUsers = messaging.empty;
        $scope.userInfoStatus = messaging.unavailable;
        console.log(status, data);
      });
    };
  });
