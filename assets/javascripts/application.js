// ---- collections ----------------------------------------------------------
var NavItemCollection = Backbone.Collection.extend({
  url: ''
});

var PostCollection = Backbone.Collection.extend({
  url: '/api/posts'
});

// ---- views ----------------------------------------------------------------
var NavSidebar = Backbone.View.extend({
  el: '.nav-sidebar',
  render: function () {
    var template = _.template($('#nav-sidebar-template').html(), {});
    this.$el.html(template);
  }
});

var PostsView = Backbone.View.extend({
  el: '.posts',
  render: function () {
    var _this = this;
    var posts = new PostCollection();
    posts.fetch({
      success: function (posts) {
        var template = _.template($('#posts-template').html(), {
          posts: posts.models
        })
        _this.$el.html(template)
      }
    });
  }
});

// ---- routes ---------------------------------------------------------------
var Router = Backbone.Router.extend({
  routes: {
    '': 'home'
  }
});

var nav = new NavSidebar();
var posts = new PostsView();

var router = new Router();

router.on('route:home', function () {
  nav.render();
  posts.render();
});

Backbone.history.start();

