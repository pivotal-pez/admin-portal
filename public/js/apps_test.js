describe('AppsController', function() {
  beforeEach(module('AppsApp'));

  var $controller;
  var $httpBackend;
  var appsInfoURI = "/v1/info/apps";

  beforeEach(inject(function($injector){
    $httpBackend = $injector.get('$httpBackend');
    $controller = $injector.get('$controller');
    $window = $injector.get('$window');
  }));

  describe('$scope.* appCounts', function() {
    it('should be initialized as undefined', function() {
      var $scope = {};
      var controller = $controller('AppsController', { $scope: $scope });
      expect($scope.totalMemory).toEqual(undefined);
      expect($scope.totalDisk).toEqual(undefined);
      expect($scope.totalInstanceCount).toEqual(undefined);
      expect($scope.totalRunningCount).toEqual(undefined);
      expect($scope.totalAppCount).toEqual(undefined);
      expect($scope.javaBPCount).toEqual(undefined);
      expect($scope.rubyBPCount).toEqual(undefined);
      expect($scope.nodeBPCount).toEqual(undefined);
      expect($scope.goBPCount).toEqual(undefined);
      expect($scope.pythonBPCount).toEqual(undefined);
      expect($scope.phpBPCount).toEqual(undefined);
      expect($scope.otherBPCount).toEqual(undefined);
      expect($scope.stoppedStateCount).toEqual(undefined);
      expect($scope.startedStateCount).toEqual(undefined);
      expect($scope.diegoAppsCount).toEqual(undefined);
    });
  });

  describe('.getAppsInfo', function() {
    controlTotalInstanceCount = 70;
    controlTotalRunningCount = 45 ;
    controlTotalMemory = 45 ;
    controlTotalDisk = 45 ;
    controlTotalAppCount = 156;
    controlJavaBPCount = 21;
    controlRubyBPCount = 8;
    controlNodeBPCount = 0;
    controlGOBPCount = 5;
    controlPythonBPCount = 0;
    controlPHPBPCount = 2;
    controlOtherBPCount = 14;
    controlStoppedStateCount = 0;
    controlStartedStateCount = 0;
    controlDiegoAppsCount = 0;
    controlEmpty = "-";

    describe("when successful response", function() {
      var $scope = {};
      it('should populate the app fields for display', function() {
        $httpBackend.when('GET', appsInfoURI).respond({"TotalInstanceCount" : controlTotalInstanceCount,
                                                      "TotalRunningCount" : controlTotalRunningCount,
                                                      "TotalMemory": controlTotalMemory,
                                                      "TotalDisk": controlTotalDisk,
                                                      "TotalAppCount" : controlTotalAppCount,
                                                      "JavaBPCount" : controlJavaBPCount,
                                                      "RubyBPCount" : controlRubyBPCount,
                                                      "NodeBPCount" : controlNodeBPCount,
                                                      "GOBPCount" : controlGOBPCount,
                                                      "PythonBPCount" : controlPythonBPCount,
                                                      "PHPBPCount" : controlPHPBPCount,
                                                      "OtherBPCount" : controlOtherBPCount,
                                                      "StoppedStateCount" : controlStoppedStateCount,
                                                      "StartedStateCount" : controlStartedStateCount,
                                                      "DiegoAppsCount" : controlDiegoAppsCount});
        var controller = $controller('AppsController', { $scope: $scope });
        controller.getAppsInfo();
        $httpBackend.flush();

        expect($scope.totalInstanceCount).toBe(controlTotalInstanceCount);
        expect($scope.totalMemory).toBe(controlTotalMemory);
        expect($scope.totalDisk).toBe(controlTotalDisk);
        expect($scope.totalRunningCount).toBe(controlTotalRunningCount);
        expect($scope.totalAppCount).toBe(controlTotalAppCount);
        expect($scope.javaBPCount).toBe(controlJavaBPCount);
        expect($scope.rubyBPCount).toBe(controlRubyBPCount);
        expect($scope.nodeBPCount).toBe(controlNodeBPCount);
        expect($scope.goBPCount).toBe(controlGOBPCount);
        expect($scope.pythonBPCount).toBe(controlPythonBPCount);
        expect($scope.phpBPCount).toBe(controlPHPBPCount);
        expect($scope.otherBPCount).toBe(controlOtherBPCount);
        expect($scope.stoppedStateCount).toBe(controlStoppedStateCount);
        expect($scope.startedStateCount).toBe(controlStartedStateCount);
        expect($scope.diegoAppsCount).toBe(controlDiegoAppsCount);
        expect($scope.userInfoStatus).toBe("");
      });
    });

    describe("when un-successful response", function() {
      var $scope = {};
      it('should populate the user fields for display', function() {
        $httpBackend.when('GET', appsInfoURI).respond(403);
        var controller = $controller('AppsController', { $scope: $scope });
        controller.getAppsInfo();
        $httpBackend.flush();

        expect($scope.totalInstanceCount).toBe(controlEmpty);
        expect($scope.totalMemory).toBe(controlEmpty);
        expect($scope.totalDisk).toBe(controlEmpty);
        expect($scope.totalRunningCount).toBe(controlEmpty);
        expect($scope.totalAppCount).toBe(controlEmpty);
        expect($scope.javaBPCount).toBe(controlEmpty);
        expect($scope.rubyBPCount).toBe(controlEmpty);
        expect($scope.nodeBPCount).toBe(controlEmpty);
        expect($scope.goBPCount).toBe(controlEmpty);
        expect($scope.pythonBPCount).toBe(controlEmpty);
        expect($scope.phpBPCount).toBe(controlEmpty);
        expect($scope.otherBPCount).toBe(controlEmpty);
        expect($scope.stoppedStateCount).toBe(controlEmpty);
        expect($scope.startedStateCount).toBe(controlEmpty);
        expect($scope.diegoAppsCount).toBe(controlEmpty);
        expect($scope.userInfoStatus).not.toBe("");
      });
    });
  });
});
