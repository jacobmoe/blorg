define([
  'jquery',
  'underscore',
  'backbone',
  'router', 
], function($, _, Backbone, Router){
  var initialize = function(){
    // Pass in our Router module and call it's initialize function
    Router.initialize();
  }

  return {
    initialize: initialize
  };
});

// ---- views ----------------------------------------------------------------
// var NavSidebar = Backbone.View.extend({
//   el: '.nav-sidebar',
//   render: function () {
//     var template = _.template($('#nav-sidebar-template').html(), {});
//     this.$el.html(template);
//   }
// });


// ---- routes ---------------------------------------------------------------
