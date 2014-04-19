require.config({
  baseUrl: "/javascripts",
  paths: {
    jquery: 'vendor/jquery-2.1.0.min',
    underscore: 'vendor/underscore-1.6.0.min',
    backbone: 'vendor/backbone-1.1.2.min',
    bootstrap: 'vendor/bootstrap.min',
    prettify: 'vendor/google-code-prettify/prettify'
  },
  urlArgs: "bust=" +  (new Date()).getTime() // don't use in prod
});

require(['application'], function(App){
  App.initialize();
});
