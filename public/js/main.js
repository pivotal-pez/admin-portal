var mainApp = angular.module('MainApp', ['AppsApp', 'UsersApp'], function($interpolateProvider) {
      $interpolateProvider.startSymbol('{*{');
      $interpolateProvider.endSymbol('}*}');
  });
