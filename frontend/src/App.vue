<template>
  <v-app dark style="background: #12171a;">
    <v-navigation-drawer
      mobile-break-point="800"
      :mini-variant="false"
      width="170"
      app
      dark
      v-model="drawer"
    >
      <v-list two-line>
        <v-list-tile v-for="(item, index) in items" :key="index" :to="item.path">
          <v-list-tile-content>
            <v-list-tile-title v-html="item.title"></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar flat fixed app dark>
      <v-toolbar-title>
        <span style="color:#fff">GO</span>
        <span style="color:#ff5722">DESKTOP</span>
      </v-toolbar-title>
    </v-toolbar>
    <v-content dark class="mb-4">
      <v-container fluid fill-height>
        <v-layout>
          <v-flex>
            <router-view></router-view>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="#263238" dark class="px-3 d-flex justify-content-between" app>
      <span>&copy; 2019 - present your_company</span>
    </v-footer>
    <flash-message class="flash-message"></flash-message>
  </v-app>
</template>

<script>
require("vue-flash-message/dist/vue-flash-message.min.css");

export default {
  data() {
    return {
      token: "",
      user: null,
      drawer: true,
      user_valid: false,
      url: "",
      busy: false,
      items: [
        {
          path: "/",
          title: "Home"
        },
        {
          path: "/settings",
          title: "Settings"
        }
      ],
      right: null
    };
  },
  mounted() {},
  methods: {}
};
</script>



<style type="text/css">
.v-navigation-drawer > .list:not(.list--dense) .list__tile {
  font-size: 17px;
}
.avatar {
  max-width: 75px;
}
.v-list__tile--link {
  border-left: 10px solid transparent;
}
.v-list__tile--link.v-list__tile--active {
  color: #ff5722 !important;
  border-left: 10px solid #ff5722;
}
.flash-message {
  display: flex;
  position: fixed;
  top: 80px;
  right: 10px;
  z-index: 999;
}
.flash__message {
  z-index: 999;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  padding: 15px 30px 15px 15px;
  margin-bottom: 10px;
  border: 1px solid transparent;
  border-radius: 4px;
  background-color: #cccccc;
  position: relative;
}
.flash__message.success {
  color: #3c763d !important;
  background-color: #dff0d8 !important;
  border-color: #d6e9c6 !important;
}
.flash__message.warning {
  color: #413115 !important;
  background-color: #fcf8e3 !important;
  border-color: #faebcc !important;
}
.flash__message.error {
  color: #a94442 !important;
  background-color: #f2dede !important;
  border-color: #ebccd1 !important;
}
.flash__message.info {
  color: #37516a !important;
  background-color: #f2f5fe !important;
  border-color: rgb(223, 227, 240) !important;
}

.v-content {
  padding: 0 20px 50px;
  background-color: #12171a !important;
}

.v-toolbar__content {
  background-color: #12171a !important;
}

.v-navigation-drawer {
  background-color: #12171a !important;
}

.loading:after {
  content: " .";
  animation: dots 1s steps(5, end) infinite;
}

@keyframes dots {
  0%,
  20% {
    color: rgba(0, 0, 0, 0);
    text-shadow: 0.25em 0 0 rgba(0, 0, 0, 0), 0.5em 0 0 rgba(0, 0, 0, 0);
  }
  40% {
    color: black;
    text-shadow: 0.25em 0 0 rgba(0, 0, 0, 0), 0.5em 0 0 rgba(0, 0, 0, 0);
  }
  60% {
    text-shadow: 0.25em 0 0 black, 0.5em 0 0 rgba(0, 0, 0, 0);
  }
  80%,
  100% {
    text-shadow: 0.25em 0 0 black, 0.5em 0 0 black;
  }
}
</style>
