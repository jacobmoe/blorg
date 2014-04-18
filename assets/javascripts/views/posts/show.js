define([
  'underscore',
  'backbone',
  'models/post'
], function(_, Backbone, Post){    
    
  var PostShow = Backbone.View.extend({
    tagName : "li",
    render : function() {
      return this;
    }
  });

  return PostShow;

}); 