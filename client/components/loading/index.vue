<template>
  <div v-if="loading" class ="loading-page">
  	<img src="~assets/loading.gif"/>
  </div>
</template>
<script >
	import EventBus from '../../assets/eventBus'
  export default {
  	name:"loading-component",
  	data(){
  		return{
  			loading:false,
  			openListener: () => {
				if (this.loading==true)return;
	            this.showing = true;
	        },
	        closeListener:() => {
	        	if (this.loading==false)return;
	            this.showing = false;
	        }
  		}
  	},
  	created(){
  		EventBus.$on("open-loading",this.openListener)
  		EventBus.$on("close-loading",this.closeListener)
  	},
  	beforeDestroy() {
        eventBus.$off('open-loading', this.openListener);
    }
  }
</script>
<style lang="scss" scoped>
.loading-page{
	z-index:10000;
	width:100vw;
	height:100vh;
	background-color:black;
}	
</style>