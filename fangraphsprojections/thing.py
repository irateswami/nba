from selenium import webdriver
from selenium.webdriver.firefox.options import Options

options = Options()
options.headless = True

browser = webdriver.Firefox(options=options)

browser.get("https://www.fangraphs.com/projections.aspx?pos=all&stats=bat&type=steamer")

element = browser.find_element_by_id('ProjectionBoard1_cmdCSV')
element.click()
