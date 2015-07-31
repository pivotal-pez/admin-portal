angular.module('AppsApp', [])
  .controller('AppsController', function($scope, $http, $interval, $timeout, $window) {
    var apps = this;
    var refreshInterval = 30000;
    var appsInfoURI = "/v1/info/apps";
    var messaging = {
      unavailable: "** Service Unavailable **",
      empty: "-"
    };

    $interval( function(){ apps.getAppsInfo(); }, refreshInterval, 0, true);

    $timeout(function () {
      apps.getAppsInfo();
    }, 1);

    // Chart.js Options
    $scope.options =  {
      // Sets the chart to be responsive
      responsive: true,
      //Boolean - Whether we should show a stroke on each segment
      segmentShowStroke : true,
      //String - The colour of each segment stroke
      segmentStrokeColor : '#fff',
      //Number - The width of each segment stroke
      segmentStrokeWidth : 2,
      //Number - The percentage of the chart that we cut out of the middle
      percentageInnerCutout : 50, // This is 0 for Pie charts
      //Number - Amount of animation steps
      animationSteps : 100,
      //String - Animation easing effect
      animationEasing : 'easeOutBounce',
      //Boolean - Whether we animate the rotation of the Doughnut
      animateRotate : true,
      //Boolean - Whether we animate scaling the Doughnut from the centre
      animateScale : false,
      //String - A legend template
      legendTemplate : '<ul class="tc-chart-js-legend"><% for (var i=0; i<segments.length; i++){%><li><span style="background-color:<%=segments[i].fillColor%>"></span><%if(segments[i].label){%><%=segments[i].label%><%}%></li><%}%></ul>'
    };

    apps.setDonutUX = function(data) {
      $scope.data = [
        { value: data.JavaBPCount, color: '#F7464A', highlight: '#FF5A5E', label: 'Java'},
        { value: data.RubyBPCount, color: '#46BFBD', highlight: '#5AD3D1', label: 'Ruby'},
        { value: data.NodeBPCount, color: '#46BFBD', highlight: '#5AD3D1', label: 'NodeJs'},
        { value: data.GOBPCount, color: '#46BFBD', highlight: '#5AD3D1', label: 'GOLang'},
        { value: data.PythonBPCount, color: '#46BFBD', highlight: '#5AD3D1', label: 'Python'},
        { value: data.PHPBPCount, color: '#FDB45C', highlight: '#FFC870', label: 'PHP'},
        { value: data.OtherBPCount, color: '#FDB45C', highlight: '#FFC870', label: 'Other'}
      ];
    }

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
        apps.setDonutUX(data);
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
