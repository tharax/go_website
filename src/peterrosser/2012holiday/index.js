var backgroundColour = "rgba(256, 256, 256, 1)";
var weekendColour    = "rgba(0, 0, 0, 0.1)";
var flightsColour    = "rgba(0, 255, 0, 0.25)";
var daysLineColour   = "rgba(0, 0, 0, 1)";
        
function draw() {
        
        var canvas = document.getElementById("NewZealandTime");
        var ctx = canvas.getContext("2d");
        var firstDayOfHoliday = new Date(2012, 12, 7);
        var lastDayOfHoliday = new Date(2013, 1, 10);
        var daysDisplayed = differenceInTime(firstDayOfHoliday, lastDayOfHoliday) + 1;
        var heightOfADay = resizeCanvas(canvas, daysDisplayed, heightOfADay);
        var widthOfADay = canvas.width / daysDisplayed;
        var hour = widthOfADay / 24;
        
        canvas.width = canvas.width + 1; //Settles drawing issues
        canvas.height = heightOfADay * 4;
        
        ctx.fillStyle = backgroundColour;
        ctx.fillRect(0, 0, canvas.width, heightOfADay);
        
        drawDays(ctx, daysDisplayed, widthOfADay, heightOfADay, firstDayOfHoliday, lastDayOfHoliday);
        drawFlights(ctx, hour, heightOfADay, widthOfADay);
}

function resizeCanvas(canvas, daysDisplayed, heightOfADay) {
        /* The idea is there is 1 pixel per minute. So if a day is only 24 pixels wide,
           it should be 60 pixels tall. The supported sizes are currently:
           48 * 30
           30 * 48
           24 * 60
           16 * 90
           12 * 120 */
        if (window.innerWidth >= (daysDisplayed * 48)) {
                canvas.width = daysDisplayed * 48;
                return 30;  
        }
        if (window.innerWidth < (daysDisplayed * 48)) {
                canvas.width = daysDisplayed * 30;
                return 48;  
        } 
        if (window.innerWidth < (daysDisplayed * 30)) {
                canvas.width = daysDisplayed * 24;
                return 60;  
        } 
        if (window.innerWidth < (daysDisplayed * 24)) {
                canvas.width = daysDisplayed * 16;
                return 90;  
        }
        if (window.innerWidth < (daysDisplayed * 16)) {
                canvas.width = daysDisplayed * 12;
                return 120;  
        }
}

function drawTimeZone(ctx, ArriveTime, TimeZoneHeight, TimeInZone, heightOfADay) {
        ctx.fillStyle = backgroundColour;
        ctx.fillRect(ArriveTime, TimeZoneHeight, TimeInZone, heightOfADay);
        
}

function drawFlights(ctx, hour, heightOfADay, widthOfADay) {
        //Auckland - Dubai (This includes a stop-over in Sydney)
        var firstDayOfHoliday = new Date(2012, 12, 7);
        drawFlight(8, +12, +12, +4, 18.40, 19.35, hour, ctx, heightOfADay, widthOfADay);
        var departureTime = new Date(2012, 12, 8, 3, 40);
        drawFlightTwo(ctx, departureTime, 1, 19.35, heightOfADay, widthOfADay, firstDayOfHoliday);

        //Dubai - Casablanca
        drawFlight(14, +12, +4, 0, 7.50, 9.00, hour, ctx, heightOfADay, widthOfADay);

        //Marrakech - Madrid
        drawFlight(28, +12, 0, +2, 14.30, 1.50, hour, ctx, heightOfADay, widthOfADay);

        //Madrid - London Heathrow
        drawFlight(28, +12, +2, +1, 20.45, 2.25, hour, ctx, heightOfADay, widthOfADay);

        //London Heathrow - Dubai
        drawFlight(31 + 2, +12, +1, +4, 13.35, 7.00, hour, ctx, heightOfADay, widthOfADay);

        //Dubai - Singapore
        drawFlight(31 + 3, +12, +4, +8, 3.25, 7.15, hour, ctx, heightOfADay, widthOfADay);
        
        //Singapore - Brisbane
        drawFlight(31 + 7, +12, +8, +10, 15.15, 7.35, hour, ctx, heightOfADay, widthOfADay);
        
        //Brisbane - Auckland
        drawFlight(31 + 8, +12, +10, +12, 8.25, 3.10, hour, ctx, heightOfADay, widthOfADay);
}

function drawFlightTwo(ctx, UTCStart, timezone, duration, heightOfADay, widthOfADay, firstDayOfHoliday) {
        var left = 0.5 + (widthOfADay /24) * ((UTCStart.getTime() - firstDayOfHoliday.getTime()) / (60 * 60 * 1000)); //Pixel to start with
        var top = 0.5;
        if (timezone === 12) {
                //doNothing
        } 
        else {
                top += heightOfADay;
        }
        timezone -= 12;
        left += (timezone * widthOfADay / 24);
        
        ctx.fillStyle = daysLineColour;
        ctx.fillText(left, 0, heightOfADay + 10);
        ctx.fillText(top, 0, heightOfADay + 20);
        ctx.fillText(widthOfADay, 100, heightOfADay +30);
        ctx.fillText(duration * (widthOfADay / 24), 0, heightOfADay + 30);
        ctx.fillText(heightOfADay, 0, heightOfADay + 40);
        
        ctx.fillStyle = flightsColour;
        ctx.fillRect(left, top, duration * (widthOfADay / 24), heightOfADay);
}

function drawFlight(date, currentTimeZone, departTimeZone, arriveTimeZone, flightStart, flightLength, hour, ctx, heightOfADay, widthOfADay) {
        date = date - 7;
        flightStart = (date * widthOfADay) + ((flightStart + currentTimeZone - departTimeZone) * hour);
        flightLength = flightLength * hour;
        ctx.fillStyle = flightsColour;
        ctx.fillRect(flightStart + 0.5, 0 + 0.5, flightLength, heightOfADay);
        var arrivalTime = flightStart + flightLength - currentTimeZone + departTimeZone - arriveTimeZone;
        //drawTimeZone(ctx, arrivalTime, heightOfADay + 30, 100, heightOfADay);
}

function drawDays(ctx, daysDisplayed, widthOfADay, heightOfADay, firstDayOfHoliday, lastDayOfHoliday) {
        var startOfDay = 0;
        var addtime = 24 * 60 * 60 * 1000; //24 hours in milliseconds
        
        for (var i = 0; i < daysDisplayed; i++) {
                startOfDay = i * widthOfADay;
                drawADay(ctx, firstDayOfHoliday, startOfDay, widthOfADay, heightOfADay);
                firstDayOfHoliday.setTime(firstDayOfHoliday.getTime() + addtime);
        }        
}
function drawADay(ctx, date, startOfDay, widthOfADay, heightOfADay) {
        //Draw border
        ctx.strokeStyle = daysLineColour;
        ctx.strokeRect(startOfDay + 0.5, 0.5, widthOfADay, heightOfADay);               
        
        //Write date
        ctx.fillStyle = daysLineColour;
        ctx.fillText(date.getDate(), startOfDay, 10);
        
        //Shade weekends
        if (date.getDay() === 2 || date.getDay() === 3) {
                ctx.fillStyle = weekendColour;
                ctx.fillRect(startOfDay + 0.5, 0.5, widthOfADay, heightOfADay);
        }
}

function differenceInTime(startTime, endTime) {
        // Set the unit values in milliseconds
        var msecPerMinute = 1000 * 60;
        var msecPerHour   = msecPerMinute * 60;
        var msecPerDay    = msecPerHour * 24;
        
        // Get the difference in milliseconds
        var interval = endTime.getTime() - startTime.getTime();
        
        return interval / msecPerDay;
        
}