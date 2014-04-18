require.config({
  baseUrl: "/javascripts",
  paths: {
    jquery: 'vendor/jquery-2.1.0.min',
    underscore: 'vendor/underscore-1.6.0.min',
    backbone: 'vendor/backbone-1.1.2.min',
    bootstrap: 'vendor/bootstrap.mins'
  },
  urlArgs: "bust=" +  (new Date()).getTime() // don't use in prod

});

require([
  // Load our app module and pass it to our definition function
  'application',
], function(App){
  // The "app" dependency is passed in as "App"
  App.initialize();
});
