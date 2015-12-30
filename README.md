#Assignment

 Our traffic robots travel around London to report on traffic conditions. Every time a robot passes close to a tube station, it assesses the traffic condition in the area, and reports it.

#Future Improvements
1. Does not seem right to loop through all the stations to see if we are near to one.  There are many algorithms that could be used here.  Even just a simple binary search on either Lat or Long would be better, and not too hard to implement given time.  
2. Not sure of code quality give new language. 
3. There is not enough test coverage.  Some test rely on manual inspection 
4. rand isn't Random (but this doesn't matter as this is not the functionality required in production)
5. Seems wrong to store the tube locations in the robot as this had limited memory.  Checking server side if a robot is near a tube station then asking the robot to report back may be better. However my understanding of channels is limited and I do not currently feel confident in implementing that. 