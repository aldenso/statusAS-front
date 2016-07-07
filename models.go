package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type page struct {
	Title string
}

// Tomlconfig struct to read toml file components.
type Tomlconfig struct {
	FrontServer FrontServerinfo
}

// FrontServerinfo struct to configure statusAS-front
type FrontServerinfo struct {
	FrontServerName string
	FrontServerPort int
	APIServerName   string
	APIServerPort   int
	APItls          bool
}

// CreateTemplate function to create a base config.toml file
func CreateTemplate() {
	template := `# Example of config Configuration
[frontserver]
frontservername = "server2.mydom.local"
frontserverport = 9000
apiservername = "server1.mydom.local"
apiserverport = 8080
apitls = true
`
	tomlfile := "config.toml"
	if _, err := os.Stat(tomlfile); err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(tomlfile)
			if err != nil {
				fmt.Println("Error creating config.toml file", err)
				os.Exit(1)
			}
			defer file.Close()
			if _, err := file.Write([]byte(template)); err != nil {
				fmt.Printf("Can't write message\n%v\n", err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("config.toml already exist in directory.")
		os.Exit(1)
	}
	fmt.Println("config.toml created.")
	os.Exit(0)
}

// CreateAppjs create app.js with config from tomlfile
func CreateAppjs() {
	appjs := `var apiurlservices = "APIURLSERVICES/api/v1/services"

var serv = new Vue({
  el: '#services',
  data: {
    services: [],
    error: false,
    operational: 0,
    degraded: 0,
    notOperational: 0
  },
  methods: {
    loadServices: function(){
      this.$http.get(apiurlservices).then(function(response){
        this.$set('services', response.json())
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
`
	var newappjs string
	if apiserverport != 0 {
		if apitls {
			newappjs = strings.Replace(appjs, "APIURLSERVICES", "https://"+apiservername+":"+strconv.Itoa(apiserverport), 1)
		} else {
			newappjs = strings.Replace(appjs, "APIURLSERVICES", "http://"+apiservername+":"+strconv.Itoa(apiserverport), 1)
		}
	} else {
		if apitls {
			newappjs = strings.Replace(appjs, "APIURLSERVICES", "https://"+apiservername, 1)
		} else {
			newappjs = strings.Replace(appjs, "APIURLSERVICES", "http://"+apiservername, 1)
		}
	}
	appjsfile := "resources/js/app.js"
	file, err := os.Create(appjsfile)
	if err != nil {
		fmt.Println("Error creating resources/js/app.js file", err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := file.Write([]byte(newappjs)); err != nil {
		fmt.Printf("Can't write\n%v\n", err)
		os.Exit(1)
	}
}
