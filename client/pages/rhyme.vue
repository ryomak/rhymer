<template>
	<div>
	 	<b-container fluid>
	 		<b-form-input v-model="sentence"
                  type="text"
                  placeholder="Enter sentence !not english"></b-form-input>
                  <b-button :disable="!canPush" @click="search()">検索</b-button>
	        <span v-for="(r,index) in rhymes" :key="index">{{r.name}}</span>
	    </b-container>
	</div>
</template>

<script>
export default {
	data(){
		return{
			sentence:"",
			rhymes:[],
		}
	},
	methods:{
		search(){
			let query = `{word(sentence:"${this.sentence}",convert_type:"normal"){name,yomi,rhyme_words}}`
		    this.$axios.post("/api/v1/rhyme",query).then(res=>{
				this.rhymes = res.data.data.word
			})
		},
		canPush(){
			return sentence.length > 1
		}
	}
}
</script>

<style lang="scss" scoped="">


</style>
