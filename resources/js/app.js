/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T21:01:59-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-04T12:15:00-04:30
*/

new Vue({
  el: '#services',
  data: {
    services: [],
    error: false,
  },
  methods: {
    loadServices: function(){
      this.$http.get('http://192.168.125.1:8080/api/v1/services').then(function(response){
        this.$set('services', response.json())
      }, function(response) {
        this.$set('error', true)
      });
    }
  },
  ready: function(){
    this.loadServices();
    setInterval(function () {
      this.loadServices();
    }.bind(this), 15000);
  }
});

new Vue({
  el: "#date",
  data: {
    date: ""
  },
  methods: {
  showtime: function() {
    var date = new Date();
    var time = date.toTimeString();
    var dd = date.getDate();
    var mm = date.getMonth()+1;//january is month 0
    var yy = date.getFullYear();
    this.date = dd+'/'+mm+'/'+yy + ' ' +time;
    }
  },
  ready: function(){
    this.showtime();

    setInterval(function () {
    this.showtime();
    }.bind(this), 30000);
  }
});
