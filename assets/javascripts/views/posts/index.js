define([
  'jquery',
  'underscore',
  'backbone',
  'prettify',
  'collections/posts_collection',
  'templates/posts/index'
], function($, _, Backbone, prettify, PostsCollection, template) {

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
          var template = _.template($('#posts-template').html(), {
            posts: posts.models
          })
          _this.$el.html(template)
          _this.afterRender();
        }
      });
    },

    afterRender: function() {
      var code = null;
      $('pre').addClass('prettyprint').each(function(idx, el){
        code = el.firstChild;
        code.innerHTML = prettify.prettyPrintOne(code.innerHTML);
      });
    }

  });

  return PostsIndex;
});
