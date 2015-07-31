angular.module('AppsApp', [])
  .controller('AppsController', function($scope, $http, $interval, $timeout, $window) {
    var apps = this;
    var refreshInterval = 30000;
    var appsInfoURI = "/v1/info/apps";
    var messaging = {
      unavailable: "** Service Unavailable **",
      empty: "-"
    };

    $interval( function(){ users.getAppsInfo(); }, refreshInterval, 0, true);

    $timeout(function () {
      apps.getAppsInfo();
    }, 1);

    apps.getAppsInfo = function() {
      var responsePromise = $http.get(appsInfoURI);
      responsePromise.success(function(data, status, headers, config) {
        $scope.totalInstanceCount = data.TotalInstanceCount;
        $scope.totalAppCount = data.TotalAppCount;
        $scope.javaBPCount = data.JavaBPCount;
        $scope.rubyBPCount = data.RubyBPCount;
        $scope.nodeBPCount = data.NodeBPCount;
        $scope.goBPCount = data.GOBPCount;
        $scope.pythonBPCount = data.PythonBPCount;
        $scope.phpBPCount = data.PHPBPCount;
        $scope.otherBPCount = data.OtherBPCount;
        $scope.stoppedStateCount = data.StoppedStateCount;
        $scope.startedStateCount = data.StartedStateCount;
        $scope.diegoAppsCount = data.DiegoAppsCount;
        $scope.userInfoStatus = "";
      });

      responsePromise.error(function(data, status, headers, config) {
        $scope.totalInstanceCount = messaging.empty;
        $scope.totalAppCount = messaging.empty;
        $scope.javaBPCount = messaging.empty;
        $scope.rubyBPCount = messaging.empty;
        $scope.nodeBPCount = messaging.empty;
        $scope.goBPCount = messaging.empty;
        $scope.pythonBPCount = messaging.empty;
        $scope.phpBPCount = messaging.empty;
        $scope.otherBPCount = messaging.empty;
        $scope.stoppedStateCount = messaging.empty;
        $scope.startedStateCount = messaging.empty;
        $scope.diegoAppsCount = messaging.empty;
        $scope.userInfoStatus = messaging.unavailable;
        console.log(status, data);
      });
    };
  });
