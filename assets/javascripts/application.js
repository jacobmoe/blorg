// views --------------------------------------------------------------------
var NavSidebar = Backbone.View.extend({
  el: '.nav-sidebar',
  render: function () {
    this.$el.html('NEW CONTENT');
  }
});

// routes -------------------------------------------------------------------
var Router = Backbone.Router.extend({
  routes: {
    '': 'home'
  }
});

var nav = new NavSidebar();

var router = new Router();

router.on('route:home', function () {
  nav.render();
});

Backbone.history.start();

