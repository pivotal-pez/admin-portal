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

    // Chart.js Options
    $scope.options =  {
      // Sets the chart to be responsive
      responsive: true,
      //Boolean - Whether the scale should start at zero, or an order of magnitude down from the lowest value
      scaleBeginAtZero : true,
      //Boolean - Whether grid lines are shown across the chart
      scaleShowGridLines : true,
      //String - Colour of the grid lines
      scaleGridLineColor : "rgba(0,0,0,.05)",
      //Number - Width of the grid lines
      scaleGridLineWidth : 1,
      //Boolean - If there is a stroke on each bar
      barShowStroke : true,
      //Number - Pixel width of the bar stroke
      barStrokeWidth : 2,
      //Number - Spacing between each of the X value sets
      barValueSpacing : 5,
      //Number - Spacing between data sets within X values
      barDatasetSpacing : 1,
      //String - A legend template
      legendTemplate : '<ul class="tc-chart-js-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].fillColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>'
    };

    users.getUsersInfo = function() {
      var responsePromise = $http.get(userInfoURI);
      responsePromise.success(function(data, status, headers, config) {
        $scope.totalUsers = data.UserCount;
        $scope.managedUsers = data.ExternalCount;
        $scope.unManagedUsers = data.UAACount;
        $scope.orphanedUsers = data.OrphanedCount;
        $scope.data = users.NewDayOverDayDataObject();
        users.SetDayOverDayDataset(data.CreateDayOverDay);
        $scope.userInfoStatus = "";
      });

      responsePromise.error(function(data, status, headers, config) {
        $scope.totalUsers = messaging.empty;
        $scope.managedUsers = messaging.empty;
        $scope.unManagedUsers = messaging.empty;
        $scope.orphanedUsers = messaging.empty;
        $scope.userInfoStatus = messaging.unavailable;
        $scope.data = {};
        console.log(status, data);
      });
    };

    users.SetDayOverDayDataset = function(dayOverDayObject) {
      var i = 0;

      for (var key in dayOverDayObject) {

        if ( users.firstWeekOfData(i) ) {
          users.addDOWLabel(key);
        }
        i++;
      }
      return dayOverDayObject;
    };

    users.firstWeekOfData = function(i) {
      return i < 7
    };

    users.addDOWLabel = function(historicDate) {
      var dow = users.GetDOW(historicDate);
      $scope.data.labels.push(dow);
    };

    users.NewDayOverDayDataObject = function() {
      return {
        labels: [],
        datasets: [
          {
            label: 'Rolling 21',
            fillColor: 'rgba(220,220,220,0.5)',
            strokeColor: 'rgba(220,220,220,0.8)',
            highlightFill: 'rgba(220,220,220,0.75)',
            highlightStroke: 'rgba(220,220,220,1)',
            data: []
          },
          {
            label: 'Rolling 14',
            fillColor: 'rgba(151,187,205,0.5)',
            strokeColor: 'rgba(151,187,205,0.8)',
            highlightFill: 'rgba(151,187,205,0.75)',
            highlightStroke: 'rgba(151,187,205,1)',
            data: []
          },
          {
            label: 'Rolling 7',
            fillColor: 'rgba(101,157,187,0.5)',
            strokeColor: 'rgba(101,157,187,0.8)',
            highlightFill: 'rgba(101,157,187,0.75)',
            highlightStroke: 'rgba(101,157,187,1)',
            data: []
          }
        ]
      };
    }

    users.GetDOW = function(dateString) {
      var d = new Date(dateString);
      var weekday = new Array(7);
      weekday[0] = "Monday";
      weekday[1] = "Tuesday";
      weekday[2] = "Wednesday";
      weekday[3] = "Thursday";
      weekday[4] = "Friday";
      weekday[5] = "Saturday";
      weekday[6] = "Sunday";
      var n = weekday[d.getDay()];
      return n;
    };
  });
