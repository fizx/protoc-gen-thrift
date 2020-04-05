// Auto-generated from example.proto ... DO NOT EDIT

include "baseplate.thrift"

  service Ranking extends baseplate.BaseplateService {
    bool is_healthy();
    
      RankingRequest Rank (1: RankingRequest request);
  }

  struct RankingRequest {
        1: Entity context;
      
        2: list<Entity> candidates;
      
        3: RequestOptions options;
  }
  

  struct Entity {
        1: string id;
      
        2: map<string, Feature> features;
      
        3: double score;
  }
  

  struct Feature {
        1: string as_string;
      
        2: i64 as_int;
      
        3: double as_float;
      
        4: bool as_bool;
      
        5: StringArray as_string_array;
      
        6: IntArray as_int_array;
        999: FeatureValue ValueSelection;
  }
  
    enum FeatureValue {
        AsString,
      
        AsInt,
      
        AsFloat,
      
        AsBool,
      
        AsStringArray,
      
        AsIntArray,
    }

  struct StringArray {
        1: list<string> value;
  }
  

  struct IntArray {
        1: list<i64> value;
  }
  

  struct RequestOptions {
        1: i32 limit;
  }
  
