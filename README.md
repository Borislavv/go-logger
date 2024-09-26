## Logrus logger wrapper.

### ENVs: 
1. **LOGGER_LEVEL** (values: info, debug, warning, error, fatal, panic)
2. **LOGGER_OUTPUT** (values: /dev/null/, stdout, stderr, filename (example: /var/log/dev.log))
3. **LOGGER_FORMAT** (values: text, json)
4. **LOGGER_LOGS_DIR** (values: any dir from root project dir., for example: var/log)
5. **LOGGER_CONTEXT_EXTRA_FIELD** (any values which will extracts from context.Context, for example: jobID, taskID)

### Usage:

    output, cancelOutput, err := logger.NewOutput()
    if err != nil {
        panic(err)
    }
    defer cancelOutput()

	lgr, lgrCancel, err := logger.NewLogrus(output)
	if err != nil {
		panic(err)
	}
	defer lgrCancel()

	lgr.InfoMsg(ctx, "hello world", logger.Fields{
        "error": "no-errors",
        "any": "additional fields", 
    })

### Note: 

An **output, cancelOutput, err := logger.NewOutput()** must be called just once per unique output, or you will see an error  
of closing already closed file while canceling the output. This happens due to two outputs refers to the same file pointer.

**Example:**

    output1, cancelOutput1, _   := logger.NewOutput("stdout")
    output2, cancelOutput2, _   := logger.NewOutput("stdout")
    cancelOutput1() // ok
    cancelOutput2() // error: closing file already closed