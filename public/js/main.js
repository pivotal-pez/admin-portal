var mainApp = angular.module('MainApp', ['tc.chartjs','AppsApp', 'UsersApp'], function($interpolateProvider) {
      $interpolateProvider.startSymbol('{*{');
      $interpolateProvider.endSymbol('}*}');
  });
