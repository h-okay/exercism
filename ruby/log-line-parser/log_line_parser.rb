class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    @line.split(":")[1].strip
  end

  def log_level
    log_level = @line.split(":")[0]
    log_level.slice(1, log_level.size - 2).downcase
  end

  def reformat 
    "#{message} (#{log_level})"
  end
end
