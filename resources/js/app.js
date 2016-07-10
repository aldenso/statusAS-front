var apiurlservices = "https://server1.mydom.local:8080/api/v1/services"

var serv = new Vue({
  el: '#services',
  data: {
    services: [],
    oldservices: [],
    error: false,
    operational: 0,
    degraded: 0,
    notOperational: 0,
    show: false
  },
  methods: {
    loadServices: function(){
      this.$http.get(apiurlservices).then(function(response){
        this.$set('services', response.json())
        if (this.oldservices.length < 1) this.$set('oldservices', response.json());
        if (JSON.stringify(this.services) !== JSON.stringify(this.oldservices)) this.$set('show', true);
        if (JSON.stringify(this.services) === JSON.stringify(this.oldservices)) this.$set('show', false);
        this.$set('oldservices', this.services)
      }, function(response) {
        this.$set('error', true)
      }).then( function() {
      var operationalCount = 0;
      var degradedCount = 0;
      var notOperationalCount = 0;
      for (var i = 0; i < this.services.length; i++){
      var current = this.services[i];
      if (current.status == 0) operationalCount++
      if (current.status == 1) degradedCount++
      if (current.status == 2) notOperationalCount++
      }
      this.$set('operational', operationalCount);
      this.$set('degraded', degradedCount);
      this.$set('notOperational', notOperationalCount);
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
