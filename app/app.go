  package app

import (
  "strconv"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

type Instruction struct {
 Id int64 `db:"id" json:"id"`
 EventStatus string `db:"event_status" json:"event_status"`
 EventName string `db:"event_name" json:"event_name"`
}

var dbmap = initDb()

func GetInstructions(c *gin.Context) {
  var instructions []Instruction
  _, err := dbmap.Select(&instructions, "SELECT * FROM instruction")
  if err == nil {
    c.JSON(200, instructions)
  } else {
    c.JSON(404, gin.H{"error": "no instruction(s) into the table"})
  }
  // curl -i http://localhost:8080/api/v1/instructions
}

func GetInstruction(c *gin.Context) {
  id := c.Params.ByName("id")
  var instruction Instruction
  
  err := dbmap.SelectOne(&instruction, "SELECT * FROM instruction WHERE id=?", id)
  if err == nil {
    instruction_id, _ := strconv.ParseInt(id, 0, 64)
  
    content := &Instruction{
      Id: instruction_id,
      EventStatus: instruction.EventStatus,
      EventName: instruction.EventName,
    }
 
    c.JSON(200, content)
  } else {
    c.JSON(404, gin.H{"error": "instruction not found"})
  }
  // curl -i http://localhost:8080/api/v1/Instructions/1
}

func PostInstruction(c *gin.Context) {
  var instruction Instruction
  c.Bind(&instruction)

  if instruction.EventStatus != "" && instruction.EventName != "" {
    if insert, _ := dbmap.Exec(`INSERT INTO instruction (event_status, event_name) VALUES (?, ?)`, instruction.EventStatus, instruction.EventName); insert != nil {
      instruction_id, err := insert.LastInsertId()
      if err == nil {
        content := &Instruction{
          Id: instruction_id,
          EventStatus: instruction.EventStatus,
          EventName: instruction.EventName,
        }
        c.JSON(201, content)
      } else {
        checkErr(err, "Insert failed")
      }
    }
  } else {
    c.JSON(422, gin.H{"error": "fields are empty"})
  }
  // curl -i -X POST -H "Content-Type: application/json" -d "{ \"event_status\": \"83\", \"event_name\": \"100\" }" http://localhost:8080/api/v1/instructions
}

func UpdateInstruction(c *gin.Context) {
  id := c.Params.ByName("id")
  var instruction Instruction
  err := dbmap.SelectOne(&instruction, "SELECT * FROM instruction WHERE id=?", id)
  
  if err == nil {
    var json Instruction
    c.Bind(&json)
    instruction_id, _ := strconv.ParseInt(id, 0, 64)
    instruction := Instruction{
      Id: instruction_id,
      EventStatus: json.EventStatus,
      EventName: json.EventName,
    }

    if instruction.EventStatus != "" && instruction.EventName != ""{
    _, err = dbmap.Update(&instruction)

      if err == nil {
        c.JSON(200, instruction)
      } else {
        checkErr(err, "Updated failed")
      }
    } else {
      c.JSON(422, gin.H{"error": "fields are empty"})
    }
  } else {
    c.JSON(404, gin.H{"error": "instruction not found"})
  }
  // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"event_status\": \"83\", \"event_name\": \"100\" }" http://localhost:8080/api/v1/instructions/1
}

func DeleteInstruction(c *gin.Context) {
  id := c.Params.ByName("id")
  var instruction Instruction
  err := dbmap.SelectOne(&instruction, "SELECT id FROM Instruction WHERE id=?", id)
  
  if err == nil {
    _, err = dbmap.Delete(&instruction)
    
    if err == nil {
      c.JSON(200, gin.H{"id #" + id: " deleted"})
    } else {
      checkErr(err, "Delete failed")
    }
  } else {
    c.JSON(404, gin.H{"error": "instruction not found"})
  }
  // curl -i -X DELETE http://localhost:8080/api/v1/instructions/1
}
