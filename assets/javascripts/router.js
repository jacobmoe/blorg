define([
  'jquery',
  'underscore',
  'backbone',
  'views/posts/index',
  'views/posts/show'
], function($, _, Backbone, PostsIndexView, PostsShowView, SidebarView) {
  
  var AppRouter = Backbone.Router.extend({
    routes: {
      ''               : 'home',
      'pages/:id/:slug': 'postsIndex',
      'posts/:id'      : 'postsShow', 
      '*actions'       : 'defaultAction'
    }
  });
  
  var initialize = function(){
    var app_router = new AppRouter;

    app_router.on('route:postsIndex', function(id) {
      console.log("POSTS INDEX ====> ", id);

      var postsIndexView = new PostsIndexView({page_id: id});
      postsIndexView.render();
    });
    
    app_router.on('route:postsShow', function(params){
      console.log("POSTS SHOW ====> ", params);

      var postsShowView = new PostsShowView();
      postsShowView.render();
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
