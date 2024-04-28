# About
Project to practice massaging, sending, and displaying large amounts of data. This project uses data about HDB resales in Singapore from 2015 - 2020. There is a total of 11k+ records.

# Features
You can search for specific towns and flat types, as well as setting a maximum price filter. The program also calculates averages for sale frequency and prices on a per year basis. Records returned are displayed in a paginated table while stats are charted on a bar graph.

# Plans
- Implement resale year and lease start year queries
- Implement more processed stats, such as top 5 cheapest and most expensive towns, trend analysis

# Dev Logs
- First Golang server, learnt the basics of http and how to send json payloads using http.
- First echo project, learnt the basics of using middleware such as limiting request body size. Echo also enables for easier request url parsing and managing Headers.
- Learnt about using query params in http requests from the frontend and how that can be integrated with zod/superforms
- First docker and mongodb project, learnt about dockerfiles, images, containers, as well as mongodb query methods 
