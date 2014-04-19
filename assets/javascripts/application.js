define([
  'jquery',
  'underscore',
  'backbone',
  'router'
], function($, _, Backbone, Router, prettify){
  var initialize = function(){
    Router.initialize();
  }

  return {
    initialize: initialize
  };
});

