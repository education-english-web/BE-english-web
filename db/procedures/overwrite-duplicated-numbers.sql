DROP PROCEDURE IF EXISTS overwriteNumber;
DELIMITER ;;

/*
overwriteNumber

This procedure receives two parameters:
- id of an office
- number of a concluded contracts

It consists the below actions within the scope of an office with the input id:
- Get the id of the latest contract with input number
- Get the latest sequence number of imported contracts in office
- Update the contract with the new sequence number

NOTE:
- This procedure is applied for numbers of imported contracts only.
- This procedure is executed in the combination with the procedure `duplicatedNumber`.
Calling this individually may produce unexpected change.
*/
CREATE PROCEDURE overwriteNumber(navisOfficeID INT, number VARCHAR(10))
BEGIN
	DECLARE latestID INT;
	DECLARE latestNumber INT;
	DECLARE newNumber INT;
	DECLARE newNumberInChar CHAR(10);
	DECLARE fullNumberInChar CHAR(10);

	SET latestID = (
		SELECT id
		FROM concluded_contracts
		WHERE concluded_contracts.navis_office_id = navisOfficeID
			AND concluded_contracts.number = number
		    AND concluded_contracts.original_contract_id IS NULL
		ORDER BY created_at DESC
		LIMIT 1
	);

	SET latestNumber = (
		SELECT MAX(CAST(TRIM(LEADING 'I' FROM concluded_contracts.number) AS UNSIGNED)) AS latest_number
			FROM concluded_contracts
			WHERE concluded_contracts.number LIKE 'I%'
			    AND concluded_contracts.navis_office_id = navisOfficeID
                AND concluded_contracts.original_contract_id IS NULL
	);

	SET newNumber = latestNumber + 1;
	SET newNumberInChar = CAST(newNumber AS CHAR(10));

	SET fullNumberInChar = CONCAT('I', LPAD(newNumberInChar, 9, '0'));

	UPDATE concluded_contracts
	SET `number` = fullNumberInChar
	WHERE id = latestID;
END;
;;
DELIMITER ;


DROP PROCEDURE IF EXISTS overwriteDuplicatedNumbers;
DELIMITER ;;

/*
overwriteDuplicatedNumbers

In an office, filter imported contracts that have the same number, update the latest one with new sequence numbers.
*/
CREATE PROCEDURE overwriteDuplicatedNumbers(navisOfficeID INT)
BEGIN
    DECLARE bDone INT;

    DECLARE n VARCHAR(10);

    -- find the list of numbers that have duplicated records, and put result into a cursor
    DECLARE curs CURSOR FOR
        SELECT `number`
        FROM concluded_contracts
        WHERE navis_office_id = navisOfficeID
            AND concluded_contracts.original_contract_id IS NULL
        GROUP BY `number`
        HAVING COUNT(id) > 1;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET bDone = 1;

    OPEN curs;

    SET bDone = 0;

    -- loop through the cursor
    REPEAT
        -- fetch each number into n
        FETCH curs INTO n;
        IF NOT bDone THEN
            -- call the function to overwrite the new sequence number
            CALL overwriteNumber(navisOfficeID, n);
        END IF;
    UNTIL bDone END REPEAT;

    CLOSE curs;
END;
;;

DELIMITER ;

