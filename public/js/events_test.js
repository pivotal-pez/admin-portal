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
        $httpBackend.when('GET', eventsInfoURI).respond([{"metadata":{"guid":"eebe7c6a-bcdd-4dcc-b824-0afec89f5b77","url":"/v2/events/eebe7c6a-bcdd-4dcc-b824-0afec89f5b77","created_at":"2015-08-07T15:55:48Z","updated_at":""},"entity":{"actee":"b85f201f-824a-44af-9488-53cffa713ec2","actee_name":"schoi-python","actee_type":"app","actor":"b85f201f-824a-44af-9488-53cffa713ec2","actor_name":"schoi-python","actor_type":"app","metadata":{"exit_description":"failed to accept connections within health check timeout","exit_status":1,"index":0,"instance":"ec5d8bbf2e64469fa05f9ae047473a22","reason":"CRASHED"},"organization_guid":"3f94d639-83cb-45a9-8ebe-cda009ac76d0","space_guid":"36430149-d7a5-4822-97ae-a74d36a13fb9","timestamp":"2015-08-07T15:55:48Z","type":"app.crash"}}]);
        var controller = $controller('EventsController', { $scope: $scope });
        controller.getEventsInfo();
        $httpBackend.flush();

        expect($scope.events[0].entity.type).toBe("app.crash");
        expect($scope.events[0].entity.actor_name).toBe("schoi-python");
        expect($scope.events[0].entity.actee_name).toBe("schoi-python");
        expect($scope.events[0].entity.metadata).toEqual({ exit_description: 'failed to accept connections within health check timeout', exit_status: 1, index: 0, instance: 'ec5d8bbf2e64469fa05f9ae047473a22', reason: 'CRASHED' });
        expect($scope.events[0].entity.timestamp).toBe("2015-08-07T15:55:48Z");
      });
    });
  });
});
