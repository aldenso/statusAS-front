/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T21:01:59-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-03T21:20:28-04:30
*/

new Vue({
  el: '#services',
  data: {
    components: [],
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
