describe('EventsController', function() {
  beforeEach(module('EventsApp'));

  var $controller;
  var $httpBackend;
  var eventsInfoURI = "/v1/info/events";

  beforeEach(inject(function($injector){
    $httpBackend = $injector.get('$httpBackend');
    $controller = $injector.get('$controller');
    $window = $injector.get('$window');
  }));

  describe('$scope.* events', function() {
    it('should be initialized as undefined', function() {
      var $scope = {};
      var controller = $controller('EventsController', { $scope: $scope });
    });
  });

  describe('.getEventsInfo', function() {

    describe("when successful response", function() {
      var $scope = {};

      it('should populate the app fields for display', function() {
        $httpBackend.when('GET', eventsInfoURI).respond({"TotalInstanceCount" : "" });
        var controller = $controller('EventsController', { $scope: $scope });
        controller.getEventsInfo();
        $httpBackend.flush();

        expect($scope.events).toBe(true);
      });
    });
  });
});
