describe('UsersController', function() {
  beforeEach(module('UsersApp'));

  var $controller;
  var $httpBackend;
  var userInfoURI = "/v1/info/users";

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
      });
    });
  });
});
