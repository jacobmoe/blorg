define([
  'underscore',
  'backbone'
], function(_, Backbone) {
  
  var Post = Backbone.Model.extend({
    title: null,
    sections: null,
    date: null
  });

  return Post;
});
