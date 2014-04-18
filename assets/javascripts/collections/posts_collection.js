define([
  'underscore',
  'backbone',
  'models/post'
], function(_, Backbone, Post) {

  var PostsCollection = Backbone.Collection.extend({
      model: Post,

      initialize: function(models, options) {
        this.options = options;
      },
      
      url: function () {
        return '/api/pages/' + this.options.page_id + '/posts';
      }
  });

  return PostsCollection;
});
