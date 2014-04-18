define([
  'jquery',
  'underscore',
  'backbone',
  'collections/posts_collection',
  'templates/posts/index'
], function($, _, Backbone, PostsCollection, template) {

  var PostsIndex = Backbone.View.extend({
    el: '.posts',

    initialize: function (params) {
      this.page_id = params.page_id;
    },

    render: function () {
      var _this = this;

      var posts = new PostsCollection([], {page_id: this.page_id});

      posts.fetch({
        success: function (posts) {
          debugger;
          var template = _.template($('#posts-template').html(), {
            posts: posts.models
          })
          _this.$el.html(template)
        }
      });
    }
  });

  return PostsIndex;
});
