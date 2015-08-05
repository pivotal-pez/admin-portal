describe('UsersController', function() {
  beforeEach(module('UsersApp'));

  var $controller;
  var $httpBackend;
  var userInfoURI = "/v1/info/users";
  var controlDayOverDaySet = {
    "2015-07-15":6,
    "2015-07-16":4,
    "2015-07-17":1,
    "2015-07-18":1,
    "2015-07-19":0,
    "2015-07-20":11,
    "2015-07-21":3,
    "2015-07-22":0,
    "2015-07-23":2,
    "2015-07-24":3,
    "2015-07-25":0,
    "2015-07-26":0,
    "2015-07-27":2,
    "2015-07-28":4,
    "2015-07-29":1,
    "2015-07-30":1,
    "2015-07-31":1,
    "2015-08-01":0,
    "2015-08-02":0,
    "2015-08-03":0,
    "2015-08-04":2
  };

  beforeEach(inject(function($injector){
    $httpBackend = $injector.get('$httpBackend');
    $controller = $injector.get('$controller');
    $window = $injector.get('$window');
  }));

  describe('$scope.* (user counts)', function() {
    it('should be initialized as undefined', function() {
      var $scope = {};
      var controller = $controller('UsersController', { $scope: $scope });
      expect($scope.totalUsers).toEqual(undefined);
      expect($scope.managedUsers).toEqual(undefined);
      expect($scope.unManagedUsers).toEqual(undefined);
      expect($scope.orphanedUsers).toEqual(undefined);
    });
  });

  describe('.getUsersInfo', function() {
    controlEmpty = "-";
    controlUserCount = 325;
    controlUAACount = 72;
    controlExternalCount = 253;
    controlOrphanedCount = 0;
    describe("when successful response", function() {
      var $scope = {};
      it('should populate the user fields for display', function() {
        $httpBackend.when('GET', userInfoURI).respond({"UserCount":controlUserCount,
                                                           "UAACount":controlUAACount,
                                                           "ExternalCount":controlExternalCount,
                                                           "OrphanedCount":controlOrphanedCount});
        var controller = $controller('UsersController', { $scope: $scope });
        controller.getUsersInfo();
        $httpBackend.flush();

        expect($scope.totalUsers).toBe(controlUserCount);
        expect($scope.managedUsers).toBe(controlExternalCount);
        expect($scope.unManagedUsers).toBe(controlUAACount);
        expect($scope.orphanedUsers).toBe(controlOrphanedCount);
        expect($scope.userInfoStatus).toBe("");
        expect($scope.data).not.toBe({});
      });
    });

    describe("when successful response on a day over day user request", function() {
      var $scope = {};
      it('should return the proper labels for the display', function() {
        $httpBackend.when('GET', userInfoURI).respond({"CreateDayOverDay":controlDayOverDaySet});
        var controller = $controller('UsersController', { $scope: $scope });
        controller.getUsersInfo();
        $httpBackend.flush();

        expect($scope.data).not.toEqual({});
        expect($scope.data.labels.length).toEqual(7);
        expect($scope.data.labels.contains("Monday")).toBe(true);
        expect($scope.data.labels.contains("Tuesday")).toBe(true);
        expect($scope.data.labels.contains("Wednesday")).toBe(true);
        expect($scope.data.labels.contains("Thursday")).toBe(true);
        expect($scope.data.labels.contains("Friday")).toBe(true);
        expect($scope.data.labels.contains("Saturday")).toBe(true);
        expect($scope.data.labels.contains("Sunday")).toBe(true);
      });

      xit('should return the proper data sets for 3 weeks of history', function() {
        $httpBackend.when('GET', userInfoURI).respond({"CreateDayOverDay":controlDayOverDaySet});
        var controller = $controller('UsersController', { $scope: $scope });
        controller.getUsersInfo();
        $httpBackend.flush();

        expect($scope.data).not.toEqual({});
        expect($scope.data.datasets[0].data).not.toEqual([]);
        expect($scope.data.datasets[1].data).not.toEqual([]);
        expect($scope.data.datasets[2].data).not.toEqual([]);
      });

    });

    describe("when un-successful response", function() {
      var $scope = {};
      it('should populate the user fields for display', function() {
        $httpBackend.when('GET', userInfoURI).respond(403);
        var controller = $controller('UsersController', { $scope: $scope });
        controller.getUsersInfo();
        $httpBackend.flush();

        expect($scope.totalUsers).toBe(controlEmpty);
        expect($scope.managedUsers).toBe(controlEmpty);
        expect($scope.unManagedUsers).toBe(controlEmpty);
        expect($scope.orphanedUsers).toBe(controlEmpty);
        expect($scope.userInfoStatus).not.toBe("");
        expect($scope.data).toEqual({});
      });
    });
  });
});

Array.prototype.contains = function(obj) {
    var i = this.length;
    while (i--) {
        if (this[i] === obj) {
            return true;
        }
    }
    return false;
}
