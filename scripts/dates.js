class DatesChooser
{
    constructor()
    {
        this.dayWindow = document.getElementById("day");
        if (this.dayWindow === null)
        {
            this.error = true;
            return;
        }
        this.dayWindow.onclick = () => {
          daysList(this);
        };

        this.monthWindow = document.getElementById("month");
        if (this.monthWindow === null)
        {
            this.error = true;
            return;
        }
        this.monthWindow.onclick = () => {
          monthList(this);
        };

        this.yearWindow = document.getElementById("year");
        if (this.yearWindow === null)
        {
            this.error = true;
            return;
        }
        this.yearWindow.onclick = () => {
            yearList(this);
        };

        this.error = false;
        this.day = Number(this.dayWindow.innerText);
        this.month = this.Months.indexOf(this.monthWindow.innerText);
        this.year = Number(this.yearWindow.innerText);
    }

    get Months()
    {
        // return ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
        return ['Января', 'Февраля', 'Марта', 'Апреля', 'Мая', 'Июня', 'Июля', 'Августа', 'Сентября', 'Октября', 'Ноября', 'Декабря'];
    }

    get DayWindow()
    {
        return this.dayWindow;
    }

    get MonthWindow()
    {
        return this.monthWindow;
    }

    get YearWindow()
    {
        return this.yearWindow;
    }

    get daysLimit()
    {
        switch (this.month)
        {
            case 0:{
                return 31;
            }
            case 1:{
                if (this.year % 4 === 0)
                {
                    return 29;
                }
                else
                {
                    return 28;
                }
            }
            case 2:{
                return 31;
            }
            case 3:{
                return 30;
            }
            case 4:{
                return 31;
            }
            case 5:{
                return 30;
            }
            case 6:{
                return 31;
            }
            case 7:{
                return 31;
            }
            case 8:{
                return 30;
            }
            case 9:{
                return 31;
            }
            case 10:{
                return 30;
            }
            case 11:{
                return 31;
            }
        }
    }

    get yearsSup()
    {
        return this.yearsInf + 2;
    }

    get yearsInf()
    {
        let date = new Date();
        return date.getFullYear();
    }

    get Error()
    {
        return this.error;
    }

    set Month(m)
    {
        this.month = this.Months.indexOf(m);
        this.monthWindow.innerText = m;
        if (Number(this.day) > Number(this.daysLimit))
        {
            this.day = Number(this.daysLimit);
            this.dayWindow.innerText = String(this.day);
        }
    }

    set Year(y)
    {
        this.year = y;
        this.yearWindow.innerText = y;
        if (Number(y) % 4 !== 0)
        {
            if (this.day > this.daysLimit)
            {
                this.day = 28;
                this.dayWindow.innerText = String(this.day);
            }
        }
    }

    set Day(d)
    {
        this.day = d;
        this.dayWindow.innerText = d;
    }
}

document.addEventListener('click', function(e) {
    let target = e.target;

    let days = document.getElementById('day');
    let months = document.getElementById('month');
    let years = document.getElementById('year');
    let isDays = false;
    if (days !== null)
    {
        isDays = target == days || days.contains(target);
    }

    let isMonths = false;
    if (months !== null)
    {
        isMonths = target == months || months.contains(target);
    }

    let isYears = false;
    if (years !== null)
    {
        isYears = target == years || years.contains(target);
    }
    if (!isDays && !isMonths && !isYears)
    {
        closeAllDatesLists();
    }
});

let lastWindow = null;

function daysList(dates)
{
    if (closeAllDatesLists() === true)
    {
        return;
    }
    if (dates.Error === true)
    {
        console.log('object construct error');
        return;
    }
    lastWindow = 'days';

    let days = dates.daysLimit;

    let list = document.createElement('ul');
    list.className = 'days-list dates-list';
    list.id = 'days-list';

    for (let i = 1; i <= days; i++ )
    {
        let li = document.createElement('li');
        li.innerText = String(i);
        li.className = 'dates-list-item';
        li.onclick = function(e){
            choseDay(li, dates);
        };
        list.append(li);
    }
    dates.dayWindow.append(list)
}

function monthList(dates)
{
    if (closeAllDatesLists() === true)
    {
        return;
    }
    if (dates.Error === true)
    {
        console.log('object construct error');
        return;
    }
    lastWindow = 'month';

    let months = dates.Months;

    let list = document.createElement('ul');
    list.className = 'month-list dates-list';
    list.id = 'month-list';

    for (let i = 0; i < months.length; i++)
    {
        let li = document.createElement('li');
        li.className = 'dates-list-item';
        li.innerText = months[i];
        li.onclick = function(e){
            choseMonth(li, dates);
        };
        list.append(li);
    }
    dates.monthWindow.append(list)
}

function yearList(dates)
{
    if (closeAllDatesLists() === true)
    {
        return;
    }
    if (dates.Error === true)
    {
        console.log('object construct error');
        return;
    }
    let list = document.createElement('ul');
    list.className = 'year-list dates-list';
    list.id = 'year-list';

    for (let i = dates.yearsInf; i <= dates.yearsSup; i++)
    {
        let li = document.createElement('li');
        li.className = 'dates-list-item';
        li.innerText = i;
        li.onclick = function(e){
            choseYear(li, dates);
        };
        list.append(li);
    }

    dates.yearWindow.append(list);
}

function closeAllDatesLists()
{
    let removed = false;
    let days = document.getElementById('days-list');
    if (days !== null)
    {
        days.remove();
        removed = true;
    }

    let months = document.getElementById('month-list');
    if (months !== null)
    {
        months.remove();
        removed = true;
    }

    let years = document.getElementById('year-list');
    if (years !== null)
    {
        years.remove();
        removed = true;
    }

    return removed;
}

function choseDay(li, dates)
{
    let day = li.innerText;
    if (dates !== null)
    {
        dates.Day = Number(day);
    }
}

function choseMonth(li, dates)
{
    let month = li.innerText;
    if (dates !== null)
    {
        dates.Month = month;
    }
}

function choseYear(li, dates)
{
    let year = li.innerText;
    if (dates !== null)
    {
        dates.Year = Number(year);
    }
}