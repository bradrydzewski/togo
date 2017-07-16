String.prototype.padStart =
    function (maxLength, fillString=' ') {
        let str = String(this);
        if (str.length >= maxLength) {
            return str;
        }

        fillString = String(fillString);
        if (fillString.length === 0) {
            fillString = ' ';
        }

        let fillLen = maxLength - str.length;
        let timesToRepeat = Math.ceil(fillLen / fillString.length);
        let truncatedStringFiller = fillString
            .repeat(timesToRepeat)
            .slice(0, fillLen);
        return truncatedStringFiller + str;
    };
