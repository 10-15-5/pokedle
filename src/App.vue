<template>
	<v-app>
		<v-main>
			<div class="container">
				<div class="header-container">
					<v-img src="/src/assets/title.png" class="title-img"></v-img>
				</div>
				<div>
					<h1 class="title">Guess todays pokemon!</h1>
				</div>
				<div>
					<v-text-field
						label="Type pokemon name"
                        outlined
						class="guess-field solo"
                                    solo-inverted

					></v-text-field>
				</div>
                
			</div>
            
		</v-main>
        
	</v-app>
</template>

<script>
</script>

<style>
.header-container {
	width: 80%;
	max-width: 250px;
}

.guess-field {
	background: white;
	width: 300px;
	margin-top: 20px;
}

.title {
	margin-top: 20px;
	color: white;
}

.title-img {
	transition: transform 0.2s;
}

.title-img:hover {
	transform: scale(1.1);
}

.container {
	display: flex;
	justify-content: center;
	align-items: center;
	flex-direction: column;
	text-align: center;
	width: 100%;
	height: 100%;
}

main {
	background-image: url("./assets/background.jpg");
	background-repeat: no-repeat;
	background-size: cover;
	background-color: black;
}
</style>
