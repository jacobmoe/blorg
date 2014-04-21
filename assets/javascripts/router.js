define([
  'jquery', 
  'underscore', 
  'backbone',
  'views/posts/index',
], function($, _, Backbone, PostsIndexView) {
  
  var AppRouter = Backbone.Router.extend({
    routes: {
      ''               : 'home',
      'pages/:id/:slug': 'postsIndex',
      '*actions'       : 'defaultAction'
    }
  });
  
  var initialize = function(){
    var app_router = new AppRouter;

    app_router.on('route:postsIndex', function(id) {
      console.log("POSTS INDEX ====> ", id);

      $('ul.sidebar li').removeClass('active');
      $('ul.sidebar li#' + id).addClass('active');

      var postsIndexView = new PostsIndexView({page_id: id});
      postsIndexView.render();
    });
    
    app_router.on('route:defaultAction', function (params) {
      console.log("DEFAULT ACTION ====> ", params);
       // We have no matching route, lets display the home page 
    });

    Backbone.history.start();
  };
  return { 
    initialize: initialize
  };

});
