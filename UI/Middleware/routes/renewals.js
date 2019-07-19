const express = require('express');
const router = express.Router();
var rp = require('request-promise');


// **** setup for specific blockchain ****


const BASE_API = "https://935A346CB20A4F699DE07CCBC39F7F1E.blockchain.ocp.oraclecloud.com:443/restproxy1/bcsgw/rest/v1/transaction/"

const chaincodename = "end2end"

var body = {
    channel: "renewals",
    chaincode: chaincodename
};


const headers = {
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
    'Authorization': 'Basic Y2xvdWQuYWRtaW46Zm9DYWxANUJlYVJpbmc='
};



// OIC address 
var OICuri = "https://integration-orasenatdpltintegration02.integration.ocp.oraclecloud.com/ic/api/integration/v1/integrations"
var OICurinoauth = "http://132.145.128.247:8011/integration/BCTOGMAIL%7C01.00.0000"

// OIC setup 
const oicheaders = {
  'Content-Type': 'application/json',
  'Access-Control-Allow-Origin': '*',
  'Authorization': 'Basic bGlzYS5qb25lczpMLmpXZWxjb21lMSo='
};






/* GET home page. */
router.post('/', function(req, res, next) {
    res.render('index', { title: 'Express' });
});





/* send test renewals listing. */
router.post('/sendTestRenewal', function(req, res, next) {

  console.log("sendTestRenewal hit");

//   need to convert res to testRenewalBody
// {
// assetid: "1",
// chaincode: "end2end",
// channel: "default",
// currentprice: "$1000",
// newprice: "$99",
// pricecap: "$10000",
// priceperiod: "1 year",
// state: "Rejected"
// }
  
// testRenewalBody = [{
//   "channel":  "renewals",
//   "chaincode":  "end2end",
//   "method":  "sendRenewal",
//   "args":  ["25","$100","$200","$300","1yr","$200","rejected"]
//   }]

// renewal := &renewal{objectType, renewalId, curPrice, piPercent, piCap, piPeriod, newPrice, state}

/// assetid, curprice, picap, piperiod, newprice, state
console.log("req.body", req.body);
  
//testRenewalBody.args.push(req.assetid)

  testRenewalBody = {
    "channel":  req.body["channel"],
    "chaincode":  req.body["chaincode"],
    "method":  "sendRenewal",
    "args":[
    req.body["assetid"], 
    req.body["currentprice"], 
    req.body["pricepercent"], 
    req.body["pricecap"],
    req.body["priceperiod"],
    req.body["newprice"], 
    req.body["state"]
        ]
    }

    console.log("testrenewalbody - ", testRenewalBody);

  const options = {
    method: 'POST',
    uri:  BASE_API+"invocation",
    headers: headers,
    body: testRenewalBody,
    json: true
  }

  console.log("options - ", options)
  
  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody)
      if (parsedBody.returnCode != 'Success') {
        console.log("failure")
        console.log(parseBody.info.proxyError)
        res.json("failure").status(500)
      } else {
        console.log('suceess')
        res.json(JSON.parse(parsedBody.result.payload))
      }
    })
    .catch(function(err) {
      // POST failed...
      res.json(err)
    });
  });



    

/* GET renewals listing. */
router.get('/renewalList', function(req, res, next) {

    console.log("---------------------in renenwal list------------------");
    body.method = "getAllRenewals"
    body.args = []

    const options = {
      method: 'POST',
      uri:  BASE_API+"query",
      headers: headers,
      body: body,
      json: true
    }
  
    rp(options)
      .then(function(parsedBody) {
        // POST succeeded...
        console.log('result-----', parsedBody)
        if (parsedBody.returnCode != 'Success') {
          console.log("failure")
          console.log(parseBody.info.proxyError)
          res.json("failure").status(500)
        } else {
          console.log('suceess')
          res.json(JSON.parse(parsedBody.result.payload))
        }
      })
      .catch(function(err) {
        // POST failed...
        res.json(err)
      });

});


/* GET renewal by id */
router.get('/getRenewal', function(req, res, next) {

    body.method = "queryRenewal"
    body.args = ["6"]
    
    console.log(body)
    const options = {
      method: 'POST',
      uri:  BASE_API+"query",
      headers: headers,
      body: body,
      json: true
    }
  
    rp(options)
      .then(function(parsedBody) {
        // POST succeeded...
        console.log('result-----', parsedBody)
        if (parsedBody.returnCode != 'Success') {
          console.log("failure")
          console.log(parseBody.info.proxyError)
          res.json("failure").status(500)
        } else {
          console.log('suceess')
          res.json(JSON.parse(parsedBody.result.payload))
        }
      })
      .catch(function(err) {
        // POST failed...
        res.json(err)
      });
    
});

/* GET renewal history by id */
router.get('/renewalHistory', function(req, res, next) {

  body.method = "queryHistoryofRenewal"
  body.args = ["6"]
  
  console.log(body)
  const options = {
    method: 'POST',
    uri:  BASE_API+"query",
    headers: headers,
    body: body,
    json: true
  }

  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody)
      if (parsedBody.returnCode != 'Success') {
        console.log("failure")
        console.log(parseBody.info.proxyError)
        res.json("failure").status(500)
      } else {
        console.log('suceess')
        res.json(JSON.parse(parsedBody.result.payload))
      }
    })
    .catch(function(err) {
      // POST failed...
      res.json(err)
    });
  
});

/* show renewals in sent state */
router.get('/renewalsByState', function(req, res, next) {

  body.method = "queryRenewalsByState"
  body.args = req.body.state    //get the state as inputted by the frontend
  
  console.log(body)
  const options = {
    method: 'GET',
    uri:  BASE_API+"query",
    headers: headers,
    body: body,
    json: true
  }

  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody)
      if (parsedBody.returnCode != 'Success') {
        console.log("failure")
        console.log(parseBody.info.proxyError)
        res.json("failure").status(500)
      } else {
        console.log('suceess')
        res.json(JSON.parse(parsedBody.result.payload))
      }
    })
    .catch(function(err) {
      // POST failed...
      res.json(err)
    });
  
});


/* show renewals in sent state */
router.get('/udpateState', function(req, res, next) {

  body.method = "respondRenewal"
  body.args = req.body.state    //get the state as inputted by the frontend
  
  console.log(body)
  const options = {
    method: 'POST',
    uri:  BASE_API+"invocation",
    headers: headers,
    body: body,
    json: true
  }

  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody)
      if (parsedBody.returnCode != 'Success') {
        console.log("failure")
        console.log(parseBody.info.proxyError)
        res.json("failure").status(500)
      } else {
        console.log('suceess')
        res.json(JSON.parse(parsedBody.result.payload))
      }
    })
    .catch(function(err) {
      // POST failed...
      res.json(err)
    });
  
});

// get middleware status
router.get('/middlewarestatus', function(req, res, next) {
  res.json({status: "middleware functioning"})
});


/* GET version of blockchain */
router.get('/getbcversion', function(req, res, next) {

  const options = {
    method: 'GET',
    uri:"https://935A346CB20A4F699DE07CCBC39F7F1E.blockchain.ocp.oraclecloud.com:443/restproxy1/bcsgw/rest/version",
    headers: headers,
    json: true
  }

  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody)
      res.json(parsedBody);
    })
    .catch(function(err) {
      console.log(err)
      // POST failed...
      res.json(err)
    });
  
});




// need to fix the oic status connector

// get oic status
router.get('/oicstatus', function(req, res, next) {
  console.log("oicstatus hit");

  const options = {
    method: 'GET',
    uri: OICuri,
    headers: oicheaders,
    json: true
  }
  console.log("options -", options);

  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody.items[1].status)
    
      res.json({"status": parsedBody.items[1].status});
    })
    .catch(function(err) {
      console.log(err)
      // POST failed...
      res.json(err)
    });

});


// get oic status
router.get('/oicstats', function(req, res, next) {
  console.log("oicstatus hit");

  const options = {
    method: 'GET',
    uri: OICuri,
    headers: oicheaders,
    json: true
  }
  console.log("options -", options);

  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody)
    
      res.json(parsedBody);
    })
    .catch(function(err) {
      console.log(err)
      // POST failed...
      res.json(err)
    });

});

//get oic status
router.get('/oicstatsnoauth', function(req, res, next) {
  console.log("oicstatus hit");

  const options = {
    method: 'GET',
    uri: OICurinoauth,
    json: true
  }
  console.log("options -", options);

  rp(options)
    .then(function(parsedBody) {
      // POST succeeded...
      console.log('result-----', parsedBody)
    
      res.json(parsedBody);
    })
    .catch(function(err) {
      console.log(err)
      // POST failed...
      res.json(err)
    });

});





// get app1 status
router.get('/app1status', function(req, res, next) {
  res.json({status:"app1 functioning"}) //will need to change to actual api
});


// get app2 status
router.get('/app2status', function(req, res, next) {
  res.json({status: "app2 functioning"}) // will need to change to actual api
});




module.exports = router;
