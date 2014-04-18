define([
  'underscore',
  'backbone'
], function(_, Backbone) {
  
  var Page = Backbone.Model.extend({
    title: null
  });

  return Page;
});
