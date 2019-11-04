<template>
  <v-container fluid>
    <v-layout>
      <v-flex>
        <p>This buttons communicate with the Go backend -></p>
        <v-btn
          @click="getPriceFor('BTC')"
          class="ma-2"
          outlined
          color="yellow black--text"
        >Load Bitcoin Price</v-btn>
        <v-btn
          @click="getPriceFor('ETH')"
          class="ma-2"
          outlined
          color="blue darken-5"
        >Load Ether Price</v-btn>

        <v-card v-if="btcPrice" class="mt-5" max-width="344">
          <v-card-text>
            <div>Bitcoin Price</div>
            <p class="display-1 text--primary">${{btcPrice.toFixed(2)}}</p>
            <p>wow, that's something!</p>
          </v-card-text>
        </v-card>

        <v-card v-if="ethPrice" class="mt-5" max-width="344">
          <v-card-text>
            <div>Ether Price</div>
            <p class="display-1 text--primary">${{ethPrice.toFixed(2)}}</p>
            <p>wow, that's something!</p>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>


<script>
export default {
  data() {
    return {
      btcPrice: null,
      ethPrice: null
    };
  },
  mounted() {
    this.handleToasts();
  },
  methods: {
    getPriceFor(symbol) {
      window.backend.Pipe.PriceFor(symbol).then(result => {
        if (symbol == "BTC") {
          this.btcPrice = result.price_usd;
        }
        if (symbol == "ETH") {
          this.ethPrice = result.price_usd;
        }
      });
    },
    handleToasts() {
      window.wails.Events.On("toast", message => {
        this.flash(message, "info");
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
