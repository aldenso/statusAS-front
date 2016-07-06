var apiurlservices = "https://server1.mydom.local:8080/api/v1/services"

new Vue({
  el: '#services',
  data: {
    services: [],
    error: false,
  },
  methods: {
    loadServices: function(){
      this.$http.get(apiurlservices).then(function(response){
        this.$set('services', response.json())
      }, function(response) {
        this.$set('error', true)
      });
    },
		isEven: function(n){
			return n % 2 == 0;
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
