# utna_food

## What is this project?

* This project is LINE Bot that manages restaurant list we want to go.

## What can we do?

### 1. Register restaurant link

* We can register restaurant link with memo.

#### How to register restaurant list?

* We can register it to send restaurant link and memo.

![Register screen](https://user-images.githubusercontent.com/28858993/132768939-f230ccc0-cdc4-4892-9bda-062a1268c2c2.jpg)

### 2. Get restaurant list that we registered

* We can get restaurant list that we registered.
	* LINE Bot shows carousel of restaurant.

#### How to get restaurant list?

* We can get it by sending get command.
	```
	get
	```
* We can access link by putting "View detail" button.

![Get screen](https://user-images.githubusercontent.com/28858993/132770909-f54d4b72-ee1a-4064-bfad-7ecace415ced.jpg)

### 3. Update restaurant link and memo

* We can update restaurant link and memo.

#### How to update restaurant link and memo?

* We can update it by sending update command.
	```
	update restaurant_id URL memo
	```
	* We can check restaurant id by sending get command.

![Update screen](https://user-images.githubusercontent.com/28858993/132770738-1f17054d-c1da-458b-964c-e995ff1a7e03.jpg)

### 4. Delete restaurant link

* We can delete restaurant link that LINE Bot manages.

#### How to delete restaurant link?

* LINE Bot gives us two method to delete it.
	1. We can delete it by sending delete command.
		```
		delete restaurant_id
		```
		* We can check restaurant id by sending get command.

![Delete screen1](https://user-images.githubusercontent.com/28858993/132770822-db159053-1c9d-4b73-b6f2-09ae0a4384bb.jpg)

2. We can delete it by choosing restaurant from carousel.
	1. When we send get command, LINE Bot returns carousel of restaurant.
	2. We can delete restaurant that we want to delete by putting delete button of carousel.

![Delete screen2](https://user-images.githubusercontent.com/28858993/132770979-2ea87704-6b9f-47cb-ab72-4380c3cb3558.jpg)
