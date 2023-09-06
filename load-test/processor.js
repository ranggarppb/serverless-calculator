function generateInput(requestParams, ctx, ee, next) {
	ctx.vars["input"] = "2 add 3"

	return next();
  }

  module.exports = {
	generateInput,
  };
