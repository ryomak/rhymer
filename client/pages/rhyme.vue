<template>
	<div>
	 	<b-container fluid>
	 		<h3>検索</h3>
	 		<h5>注意:英語は使用できません・２文字以上３０文字以内</h5>
	 		<b-form-input v-model="sentence"
                  type="text"
                  placeholder="Enter sentence !not english"></b-form-input>
                  <b-button :disable="!canPush" @click="search()">検索</b-button>
	        <div v-for="(r,index) in rhymes" :key="index">
	        	{{r.name}}<p><span v-for="(w,i) in r.rhyme_words" :key="i">・{{w}}</span></p>
	        </div>
	    </b-container>
	</div>
</template>

<script>
import EventBus from '../assets/eventBus'
export default {
	data(){
		return{
			sentence:"",
			rhymes:[],
		}
	},
	methods:{
		search(){
			EventBus.$emit("open-loading")
			let query = `{word(sentence:"${this.sentence}",convert_type:"normal"){name,yomi,rhyme_words}}`
		    this.$axios.post("/api/v1/rhyme",query).then(res=>{
				this.rhymes = res.data.data.word;
				EventBus.$emit("close-loading");
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
