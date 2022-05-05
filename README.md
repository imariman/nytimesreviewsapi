# NY Times Movie Reviews API

This system is desgined to allow users to search for movie reviews over New York Times API.

**Search Parameters:** <br/>
title: Search by movie title. <br />
reviewer: Search by rewiever name. <br />
publication-date: Search by review publication date. <br />

***The system does not make a direct real-time query to the New York Times API for movies that are published after 2020.***

**System Design:**

Controller : Controller is the entry point of our application. Controller interacts with the service layer by requesting a model data. This part does not hold any logic.

Service: Service is in charge of interacting with the repository. An abstraction layer between controller and repository. It will be easy to add/change data source with this layer.

Repository: 
Repository is in charge of making HTTP calls to NY Times API and get the data.
