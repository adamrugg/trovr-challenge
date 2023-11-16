# Trovr sheepdog challenge - Adam Ruggier

The below will discuss a number of choices made throughout development, particularly on how both herding (sorting) and searching were optimised with a view towards scalability. All code changes are commented throughout the document.

## Handling storing of sheep / Herding (quickSort)

The use of quickSort was implemented as a more efficient alternative to Bubble Sort (the current herding implementation) providing a faster average time complexity of O(n log n), as opposed to Bubble Sort's average time complexity of O(n^2), handicapping its use with larger data sets, thus making quickSort a more scalable option especially as the number of sheep grow.

## Searching for sheep (Binary Search)

Binary Search was preferred due to its nature of requiring data to be sorted first, whilst also providing a more efficient option wih larger data sets with a time complexity of O(log n); the previously implemented linear search would simply take too long for larger data sets as the game progresses due to its' sequential nature (O(n)).

## General improvements (unimplemented code)

Implemented incomplete and incorrect code throughout the codebase, particularly within default_dog_impl.go (e.g. implementing Bark, a Work function) as well as including the new search and sort functions within the test and benchmark classes.

Thank you for taking the time to go through my code :\) 

# Welcome to the Trovr developer competency test

## Scenario

You are a new member of the development team at indie games development company Simulate Games. The company is developing a new text based adventure game, based around farming. You are tasked with making the game more efficient. You need to use your unique skill sets of design patterns, algorithms and data structures to improve the lackluster performance of the incomplete game.
The main character in the game is a German Shepherd Dog called Monty and his owner, farmer Ted Sneed. Farmer Ted owns a farm called Sneed's Farm, which has a single field of sheep. The game explores the adventures and the danger of owning a farm. The game player needs to help farmer Ted and his trusty companion Monty, thrive and navigate the challenging and sometimes dangerous world of farming. The player will have the opportunity to expand their farming empire as the game progresses, expanding to new fields and farms.

## Your Task

Your task as a developer is to improve how the game handles sheep and sheepdogs. You should improve the current implementation and suggest a new more efficient way of storing sheep.
You should apply all the best developer practices and comment where the current codebase diverges from these in the area you are looking at.
The code base provided is written in Go and there is a benchmark and test within the Farm Package that you will be assessed on. You will also be assessed on the code that you produce as part of this exercise. The Farm Pack is commented but not to an exhaustive level as we want you to be able to demonstrate your ability to interpret incomplete code, fill in the gaps, and solve problems as part of your improvements. We expect candidates to write their conventional commits using best practice. The time limit for this exercise is 2 hours.

## Game Rules

- A field can only have one type of livestock in it at a time.
- A field can be used to grow crops or as graze land.
- A sheep dog can herd the sheep as long as it is healthy, happy, has enough energy, and is not hungry.
- Sheep are livestock and need to be kept in a field.
- When a sheepdog first herds some sheep they are sorted by their IDs from smallest to largest.
- A player needs to be able to find sheep by their IDs using the sheepdog.
- When searching for sheep, they first must be herded or in order.
- When sheep are left for too long they get unsorted and need re-herding.
- Sheep can be attacked by predators at random interviews. If a sheep dog fails to guard the sheep they will take damage.

Please email josh@trovrapp.com, jay@trovrapp.com to confirm completion of test.

</br>
</br>
</br>
<h1>Good luck!</h1>
