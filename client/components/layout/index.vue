<template>
	<div>
		<header-component></header-component>
			<slot></slot>
		
	</div>
</template>
<script>
	import HeaderComponent from "./HeaderComponent.vue"
	import FooterComponent from "./FooterComponent.vue"
  	export default {
		name:"lay",
		components:{
			HeaderComponent,
			FooterComponent,
		}
	}
</script>