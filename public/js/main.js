var mainApp = angular.module('MainApp', ['tc.chartjs','AppsApp', 'UsersApp', 'EventsApp'], function($interpolateProvider) {
      $interpolateProvider.startSymbol('{*{');
      $interpolateProvider.endSymbol('}*}');
  });
