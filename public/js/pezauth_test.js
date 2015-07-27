describe('PezPortalController', function() {
  beforeEach(module('pezPortal'));

  var $controller;
  var $httpBackend;
  var testEmailAddy = 'test@pivotal.io';
  var testName = "Testy Larue";
  var testAPIKey = "12345";

  beforeEach(inject(function($injector){
    $httpBackend = $injector.get('$httpBackend');
    $controller = $injector.get('$controller');
    $window = $injector.get('$window');    
  }));  

  describe('$scope.myName & myEmail', function() {
    it('should be initialized as undefined', function() {
      var $scope = {};
      var controller = $controller('PezPortalController', { $scope: $scope });
      expect($scope.myName).toEqual(undefined);
      expect($scope.myEmail).toEqual(undefined);
    });
  });
  
  describe('$scope.myName & myEmail', function() {
    it('should allow initialization', function() {
      var $scope = {"myEmail": testEmailAddy, "myName": testName};
      var controller = $controller('PezPortalController', { $scope: $scope });
      expect($scope.myName).toEqual(testName);
      expect($scope.myEmail).toEqual(testEmailAddy);
    });
  });  
  
  describe('$scope.hideCLIExample', function() {
    it('should initialize to true', function() {
      var $scope = {};
      var controller = $controller('PezPortalController', { $scope: $scope });
      expect($scope.hideCLIExample).toEqual(true);      
    });
  });

  describe('pezportal.getOrgRestUri', function() {
    it('should combine the ORG API base with email into a path', function() {
      var $scope = {"myEmail": testEmailAddy};
      var controller = $controller('PezPortalController', { $scope: $scope });
      expect(controller.getOrgRestUri()).toEqual('/v1/org/user/' + testEmailAddy);
    });
  });
  
  describe('pezportal.getRestUri', function() {
    it('should combine the API base with email', function() {
      var $scope = {"myEmail": testEmailAddy};
      var controller = $controller('PezPortalController', { $scope: $scope });
      expect(controller.getRestUri()).toEqual('/v1/auth/api-key/' + testEmailAddy);
    });
  });   
  
  describe('pezportal.create', function() {
    it('should create an API Key', function() {    
      $httpBackend.when('PUT', ['/v1/auth/api-key', testEmailAddy].join('/')).respond({"APIKey": testAPIKey});
      
      var $scope = {"myEmail": testEmailAddy};
      var controller = $controller('PezPortalController', { $scope: $scope });
      controller.create();
      $httpBackend.flush();

      expect($scope.myData.APIKey).toBe(testAPIKey);
      });
  });
  
  describe('pezportal.createorg', function() {
    it('should create an createorg when orgButtonText is "Create Your Org Now"', function() {    
      $httpBackend.when('PUT', ['/v1/org/user', testEmailAddy].join('/')).respond(201);
      
      var $scope = {"myEmail": testEmailAddy, "orgButtonText": "Create Your Org Now"};
      var controller = $controller('PezPortalController', { $scope: $scope });
      controller.createorg();      
      $httpBackend.flush();

      expect($scope.orgButtonText).toBe("View Org Now");
    });
  });      
  
  describe('pezportal.createorg', function() {
    it('should not create an org when a user can\'t be found', function() {    
      $httpBackend.when('PUT', ['/v1/org/user', testEmailAddy].join('/')).respond(403, {"ErrorMsg": "query failed. unable to find matching user guid."});
      
      var $scope = {"myEmail": testEmailAddy, "orgButtonText": "Create Your Org Now"};
      var controller = $controller('PezPortalController', { $scope: $scope });
      controller.createorg();      
      $httpBackend.flush();

      expect($scope.orgButtonText).toBe("Get Okta Tile for HeritageCF");
    });
  });       
});
